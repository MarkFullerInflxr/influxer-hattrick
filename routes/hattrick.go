package hattrick

import (
	"fmt"
	list_util "influxer/hattrick/utilities"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	running_schools []string
	mutex           sync.Mutex
)

type InitiateHattrickRequest struct {
	School []string `json:school`
}
type InitiateHattrickResponse struct {
	Blocked []string
	Running []string
}

func markSchoolAsStarted(school string) {
	mutex.Lock()
	running_schools = append(running_schools, school)
	mutex.Unlock()
}

func markSchoolAsFinished(school string) {
	mutex.Lock()
	running_schools = list_util.RemoveValue(running_schools, school)
	mutex.Unlock()
}

func runHattrick(school string) {
	defer markSchoolAsFinished(school)
	fmt.Println("Running: " + school)
	time.Sleep(10 * time.Second)
	fmt.Println("Done With: " + school)
}

// @BasePath /api/v1
// Initiate Hattrick godoc
// @Summary Run Hattrick
// @Schemes InitiateHattrickRequest
// @Description run hattrick
// @Tags hattrick
// @Param requestBody body InitiateHattrickRequest true "List of Schools to Run"
// @Accept json
// @Produce json
// @Success 200 {object} InitiateHattrickResponse
// @Router /hattrick/run [post]
func Hattrick(g *gin.Context) {
	var requestData InitiateHattrickRequest

	if err := g.ShouldBindJSON(&requestData); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(requestData.School, running_schools)

	blockedSchools := []string{}
	approvedSchools := []string{}

	for _, s := range requestData.School {
		if list_util.Contains(running_schools, s) {
			blockedSchools = append(blockedSchools, s)
		} else {
			approvedSchools = append(approvedSchools, s)
		}
	}

	fmt.Println("Blocked:", blockedSchools)
	fmt.Println("Allowed:", approvedSchools)

	for _, s := range approvedSchools {
		go runHattrick(s)
		markSchoolAsStarted(s)
	}

	returnVal := InitiateHattrickResponse{
		Blocked: blockedSchools,
		Running: running_schools,
	}

	g.JSON(http.StatusOK, returnVal)
}
