// In The Name Of God

package main

import (
	"fmt"
	"todo-app/config"
	"todo-app/models"
	"todo-app/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	//Connected to DB
	config.ConnectDB()

	//Migrate Models
	config.DB.AutoMigrate(&models.Todo{})

	//Create Engin/router Gin 
	router:= gin.Default()

	//Define a route
	routes.TodoRoutes(router)

	//start web server on port 8000
	fmt.Println("Start web server...")
	router.Run(":8000")
}
