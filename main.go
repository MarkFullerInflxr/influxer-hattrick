package main

import (
	"fmt"
	docs "influxer/hattrick/docs"
	filesystem "influxer/hattrick/routes/files"
	hattrick "influxer/hattrick/routes/hattrick"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func middleware(g *gin.Context) {
	fmt.Println("Middleware!")
}

const API_BASE string = "/api/v1"

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = API_BASE // global route base path, configure it in swagger
	v1 := r.Group(API_BASE)              // global root base path, configure for gin
	{
		eg := v1.Group("/example", middleware) // begin group of api endpoints with the same base path
		{
			eg.GET("/helloworld", Helloworld) // configure single endpoint within group
		}
		hg := v1.Group("/hattrick")
		{
			hg.POST("/run", hattrick.Hattrick)
		}
		fg := v1.Group("/filesystem")
		{
			fg.GET("/", filesystem.GetFile)
		}

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
