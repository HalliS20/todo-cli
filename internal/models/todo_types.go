package models

type TodoTypes interface {
	Todo | TodoList
}

type TodoPointers interface {
	*Todo | *TodoList
}

type TodoTypeList interface {
	[]Todo | []TodoList
}

type TodoTypeListPointers interface {
	*[]Todo | *[]TodoList
}
