package controllers

import (
	"net/http"
	"todo/src/config"
	"todo/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

type todoRequest struct {
	Name        string `json:name`
	Description string `json:description`
}

type todoResponse struct {
	todoRequest
	ID uint `json:id`
}

func CreateTodo(context *gin.Context) {
	var data todoRequest

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}

	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Create(&todo)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response todoResponse

	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)

}

func AllTodos(context *gin.Context) {
	var todos []models.Todo

	err := db.Find(*&todos)

	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return

	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})

}

func UpdateTodo(context *gin.Context) {

}

func DeleteTodo(context *gin.Context) {

}
