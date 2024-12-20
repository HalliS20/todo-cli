package gormBase

import (
	"fmt"
	"reflect"
	mo "todo-cli/internal/models"

	"gorm.io/gorm"
)

// it may not work to have item here as a pointer
// since we need to check which TodoType it is and pointer may mess it up
func Create[T mo.TodoPointers](item T, db *gorm.DB) {
	result := db.Create(item)
	if result.Error != nil {
		typeName := reflect.TypeOf(item).Elem().Name()
		fmt.Println("failed to create", typeName)
		panic(result.Error)
	}
}

func Delete[T mo.TodoTypes](id int, db *gorm.DB) {
	var zero T
	result := db.Delete(&zero, id)
	if result.Error != nil {
		typeName := reflect.TypeOf(zero).Name()
		fmt.Println("failed to delete", typeName)
		panic(result.Error)
	}
}

func GetOne[T mo.TodoTypes](id int, db *gorm.DB) T {
	var item T
	result := db.Where("id = ?", id).First(item)
	if result.Error != nil {
		typeName := reflect.TypeOf(item).Elem().Name()
		fmt.Println("failed to get item", typeName)
		panic(result.Error)
	}
	return item
}

func GetAll[T mo.TodoTypes](db *gorm.DB) []T {
	var items []T
	result := db.Find(&items)
	if result.Error != nil {
		typeName := reflect.TypeOf(items).Elem().Name()
		fmt.Println("failed to get item", typeName)
		panic(result.Error)
	}
	return items
}

func UpdateOrAdd[T mo.TodoPointers](items []T, db *gorm.DB) {
	error := db.Transaction(func(tx *gorm.DB) error {
		for i := range items {
			id := reflect.ValueOf(items[i]).Elem().FieldByName("ID").Interface().(uint)
			if id == 0 {
				// WE Create an item if it has no id which indicates it doesn't exist
				err := tx.Create(items[i]).Error
				if err != nil {
					fmt.Println("failed to Create")
					panic(err)
				}
				continue
			}
			// otherwise WE just update the index since all items passed in need an index update or creation
			err := UpdateFieldTx(items[i], "Index", tx)
			if err != nil {
				fmt.Println("failed to update")
				panic(err)
			}
		}
		return nil
	})
	if error != nil {
		fmt.Println("failed to fix and add")
		panic(error)
	}
}
