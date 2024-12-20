package listRepo

import (
	"todo-cli/internal/models"

	"gorm.io/gorm"
)

type (
	TodoList = models.TodoList
)

type GormListRepo struct {
	db *gorm.DB
}

func NewListRepo(db *gorm.DB) GormListRepo {
	if db.AutoMigrate(&TodoList{}) != nil {
		panic("failed to auto migrate todos")
	}
	return GormListRepo{db}
}
