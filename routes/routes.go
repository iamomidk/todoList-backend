// routes/routes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"todoList/controllers"
	"todoList/repositories"
)

// RegisterRoutes registers application routes and their handlers.
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// Create instances of repositories and controllers
	todoRepo := repositories.NewTodoRepository(db)
	todoController := controllers.NewTodoController(todoRepo)

	// Define routes and their corresponding handlers
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to the Todo List Web Application!")
	})

	r.GET("/todos", todoController.GetAllTodos)
	r.POST("/todos", todoController.CreateTodo)
	r.PUT("/todos/:id", todoController.UpdateTodo)
	r.DELETE("/todos/:id", todoController.DeleteTodo)
}
