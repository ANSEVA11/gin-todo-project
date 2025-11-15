//In The Name Of God

package controllers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTodo(context *gin.Context) {
	var todo models.Todo

	// Parse the JSON request body into the Todo struct
	err := context.ShouldBindJSON(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the new Todo item into the database
	result := config.DB.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Send a JSON response with the created Todo item
	context.JSON(http.StatusCreated, todo)
}

func GetTodos(context *gin.Context) {
	var todos []models.Todo

	// Find all todos from database
	result := config.DB.Find(&todos)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Send the list of todos
	context.JSON(http.StatusOK, todos)
}

func GetTodo(context *gin.Context) {
	var todo models.Todo

	id := context.Param("id")

	// Find the todo by ID
	result := config.DB.First(&todo, id)

	// Case 1: Record not found
	if result.Error == gorm.ErrRecordNotFound {
		context.JSON(http.StatusNotFound, gin.H{"error": "Todo Not Found!"})
		return

		// Case 2: Some other error occurred
	} else if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Send the todo
	context.JSON(http.StatusOK, todo)
}

func UpdateTodo(context *gin.Context){
	var todo models.Todo

	id := context.Param("id")

	result := config.DB.First(&todo, id)
	if result.Error == gorm.ErrRecordNotFound {
		context.JSON(http.StatusNotFound, gin.H{"error" : "Todo Not Found!"})
		return
	}

	var new_todo models.Todo

	err := context.ShouldBind(&new_todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	todo.Title = new_todo.Title
	todo.Description = new_todo.Description
	todo.Completed = new_todo.Completed
	// Or we can use this if we have some fields:  config.DB.Model(&todo).Updates(new_todo)

	// Save changes
	result = config.DB.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error" : result.Error.Error()})
		return
	}

	// Send the todo
	context.JSON(http.StatusOK, todo)

}

func DeleteTodo(context *gin.Context){
	var todo models.Todo

	id := context.Param("id")

	result := config.DB.First(&todo, id)
	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"message" : "Todo Not Found!"})
		return
	}

	result = config.DB.Delete(&todo)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error" : result.Error.Error()})
		return
	}

	// Send Message the todo deleted successfully
	context.JSON(http.StatusOK, gin.H{"error":"Todo deleted successfully"})
}
