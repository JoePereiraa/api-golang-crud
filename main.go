package main

import (
	"log"

	"github.com/JoePereiraa/first-crud-golang/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
