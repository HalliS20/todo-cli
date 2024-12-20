package gormRepository

import (
	"todo-cli/internal/repository/gormRepository/list"
	"todo-cli/internal/repository/gormRepository/todo"

	"gorm.io/gorm"
)

type GormRepository struct {
	Todos todoRepo.GormTodoRepo
	Lists listRepo.GormListRepo
}

func NewGormRepository(db *gorm.DB) GormRepository {
	newGormRepo := GormRepository{todoRepo.NewTodoRepo(db), listRepo.NewListRepo(db)}
	return newGormRepo
}
