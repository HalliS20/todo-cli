package gormRepository

import (
	"todo-cli/internal/repository/gormRepository/todo"

	"gorm.io/gorm"
)

type GormRepository struct {
	Todos todoRepo.GormTodoRepo
}

func NewGormRepository(db *gorm.DB) GormRepository {
	newGormRepo := GormRepository{todoRepo.NewTodoRepo(db)}
	return newGormRepo
}
