package utils

import (
	"reflect"
	mo "todo-cli/internal/models/Unison"
)

func GetIndex[T mo.TodoTypes](item *T) *int {
	field := reflect.ValueOf(item).Elem().FieldByName("Index").Addr().Interface().(*int)
	return field
}
