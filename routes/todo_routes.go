// In The Name Of God

package routes

import (
	"todo-app/controllers"
	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine){

	router.POST("/api/v1/create_todo", controllers.CreateTodo)
	router.GET("/api/v1/get_todos", controllers.GetTodos)
	router.GET("/api/v1/get_todo/:id", controllers.GetTodo)
	router.PUT("/api/v1/update_todo/:id", controllers.UpdateTodo)
	router.DELETE("/api/v1/delete_todo/:id", controllers.DeleteTodo)
}