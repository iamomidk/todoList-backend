// repositories/todo_repository.go

package repositories

import (
	"gorm.io/gorm"
	"todoList/models"
)

// TodoRepository provides database interactions for the Todo model.
type TodoRepository struct {
	db *gorm.DB
}

// NewTodoRepository creates a new instance of TodoRepository.
func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

// GetAllTodos retrieves all Todos from the database.
func (r *TodoRepository) GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

// CreateTodo creates a new Todo in the database.
func (r *TodoRepository) CreateTodo(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

// UpdateTodo updates an existing Todo in the database.
func (r *TodoRepository) UpdateTodo(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

// DeleteTodoByID deletes a Todo by ID from the database.
func (r *TodoRepository) DeleteTodoByID(todoID uint) error {
	return r.db.Delete(&models.Todo{}, todoID).Error
}
