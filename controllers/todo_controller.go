// controllers/todo_controller.go

package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"todoList/models"
	"todoList/repositories"
)

// TodoController handles HTTP requests related to Todo operations.
type TodoController struct {
	repo repositories.TodoRepository
}

// NewTodoController creates a new instance of TodoController.
func NewTodoController(repo *repositories.TodoRepository) *TodoController {
	return &TodoController{
		repo: *repo,
	}
}

// GetAllTodos handles the GET request to retrieve all Todos.
func (c *TodoController) GetAllTodos(ctx *gin.Context) {
	todos, err := c.repo.GetAllTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todos"})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

// CreateTodo handles the POST request to create a new Todo.
func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := c.repo.CreateTodo(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	ctx.JSON(http.StatusCreated, todo)
}

// UpdateTodo handles the PUT request to update an existing Todo.
func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	todo.ID = uint(todoID)
	if err := c.repo.UpdateTodo(&todo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

// DeleteTodo handles the DELETE request to delete a Todo by ID.
func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	if err := c.repo.DeleteTodoByID(uint(todoID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
