package todoRepo

import (
	"todo-cli/internal/models"

	"gorm.io/gorm"
)

type (
	Todo = models.Todo
)

type GormTodoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) GormTodoRepo {
	if db.AutoMigrate(&Todo{}) != nil {
		panic("failed to auto migrate todos")
	}
	return GormTodoRepo{db}
}
