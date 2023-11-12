package files

import (
	http_util "influxer/hattrick/utilities"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type filesystemRequest struct {
	Filename string `json:filename`
}

func GetFile(g *gin.Context) {

	var requestData filesystemRequest

	if err := g.ShouldBindJSON(&requestData); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content, err := os.ReadFile(requestData.Filename)

	if err != nil {
		g.JSON(http.StatusBadRequest, "unable to open file")
	}

	mime, err := http_util.GetMimeType(requestData.Filename)

	if err != nil {
		g.JSON(http.StatusBadRequest, "unable to get mime type of file")
		return
	}
	g.Header("Content-Type", mime)
	g.String(http.StatusOK, string(content))
}
