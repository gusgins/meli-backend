package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gusgins/meli-backend/config"
	"github.com/gusgins/meli-backend/repository/mysql"
	"github.com/gusgins/meli-backend/service"
)

func main() {
	config := config.NewConfig()
	repository, err := mysql.NewRepository(config)
	if err != nil {
		fmt.Println("Could not connect to repository")
		os.Exit(1)
	}
	defer repository.Close()
	service := service.NewService(config, repository)
	r := SetupRouter(config, service)
	r.Run(fmt.Sprintf(":%d", config.API.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// SetupRouter sets routes and additional configuration
func SetupRouter(config config.Configuration, service service.Service) *gin.Engine {
	r := gin.Default()
	r.POST("/mutant", service.PostMutant)
	r.GET("/stats", service.GetStats)
	return r
}
