package routes

import (
	"todo/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	route.POST("/todos", controllers.CreateTodo)
	route.GET("/todos", controllers.AllTodos)
	route.PUT("/todos/:idTodo", controllers.UpdateTodo)
	route.DELETE("/todos/:idTodo", controllers.DeleteTodo)

	route.Run()
}
