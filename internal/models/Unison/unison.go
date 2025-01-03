package Unison

import (
	"todo-cli/internal/models/lists"
	"todo-cli/internal/models/todo"
)

type TodoTypes interface {
	todo.Todo | lists.TodoList
}

type TodoPointers interface {
	*todo.Todo | *lists.TodoList
}

type TodoTypeList interface {
	[]todo.Todo | []lists.TodoList
}

type TodoTypeListPointers interface {
	*[]todo.Todo | *[]lists.TodoList
}
