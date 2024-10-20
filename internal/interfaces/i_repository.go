package interfaces

import (
	"todo-cli/internal/models"
)

type Todo = models.Todo

type IRepository interface {
	Create(todo *Todo)
	Update(todo *Todo)
	Delete(id int)
	GetAll() []Todo
	GetOne(id int) Todo
}
