package main

import (
	"api/zeus/database"
	"api/zeus/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const port string = ":8800"

func main() {

	postgresClient := database.ConnectionPostgres()
	defer postgresClient.Close()

	router := gin.Default()
	routes.InitializeRoutes(router)

	if err := router.Run(port); err != nil {
		err := fmt.Errorf("Could not run the application: %v", err)
		log.Fatalf(err.Error())
	}
}
