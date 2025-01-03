package listRepo

import (
	"todo-cli/internal/models/lists"

	"gorm.io/gorm"
)

type (
	TodoList = lists.TodoList
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
