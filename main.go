package main

import (
	"log"

	"github.com/anasfirly20/go-codevo/controller"
	"github.com/anasfirly20/go-codevo/initializers"
	"github.com/gin-gonic/gin"
)

var (server *gin.Engine)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	initializers.ConnectDB(&config)
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	
	router := server.Group("/api")
	router.GET("/healthchecker", controller.GETUser)

	log.Fatal(server.Run(":" + config.ServerPort))
	
}