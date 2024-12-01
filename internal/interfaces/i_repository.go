package interfaces

import (
	"todo-cli/internal/models"

	"gorm.io/gorm"
)

type Todo = models.Todo

type IRepository interface {
	Create(todo *Todo)
	Update(todo *Todo)
	BatchUpdate(todos []Todo)
	UpdateField(todo Todo, field string) error
	UpdateFieldTx(db *gorm.DB, todo Todo, field string) error
	FixAndAdd(todos []Todo)
	Delete(id int)
	GetAll() []Todo
	GetOne(id int) Todo
}
