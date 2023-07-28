package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"

	"gorm.io/driver/postgres"

	"todoList/config"
	"todoList/controllers"
	"todoList/models"
	"todoList/repositories"
)

func main() {
	// Load application configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Create a GORM database connection
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// AutoMigrate the Todo model to create the table in the database
	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal("Error migrating the database connection:", err)
		return
	}

	// Close the database connection when the main function ends
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error getting the database connection:", err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Fatal("Error closing the database connection:", err)
		}
	}(sqlDB)

	// Create repositories and controllers
	todoRepo := repositories.NewTodoRepository(db)
	todoController := controllers.NewTodoController(todoRepo)

	// Create a new Gin router
	router := gin.Default()

	// Define routes and their corresponding handlers
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to the Todo List Web Application!")
	})
	router.GET("/todos", todoController.GetAllTodos)
	router.POST("/todos", todoController.CreateTodo)
	router.PUT("/todos/:id", todoController.UpdateTodo)
	router.DELETE("/todos/:id", todoController.DeleteTodo)

	// Start the server on port 8080
	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
