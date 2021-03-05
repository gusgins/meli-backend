package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gusgins/meli-backend/config"
	"github.com/gusgins/meli-backend/service"
)

func main() {
	config := config.NewConfig()
	service := service.NewService(config)
	r := SetupRouter(config, service)
	r.Run(fmt.Sprintf(":%d", config.API.APIPort)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// SetupRouter sets routes and additional configuration
func SetupRouter(config config.Configuration, service service.Service) *gin.Engine {
	r := gin.Default()
	r.POST("/mutant", service.PostMutant)
	return r
}
