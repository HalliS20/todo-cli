package gormBase

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
	mo "todo-cli/internal/models"
	fs "todo-cli/pkg/functions"
)

func UpdateFieldTx[T mo.TodoPointers](item T, field string, db *gorm.DB) error {
	fieldVal := reflect.ValueOf(item).Elem().FieldByName(field)
	if !fieldVal.IsValid() {
		return fmt.Errorf("field %s does not exist", field)
	}
	var zero T
	id := reflect.ValueOf(item).Elem().FieldByName("ID").Interface().(uint)
	if id == 0 {
		return fmt.Errorf("id not set")
	}

	return db.Model(&zero).
		Where("id = ?", id).
		Update(strings.ToLower(field), fieldVal.Interface()).
		Error
}

func BatchSave[T mo.TodoPointers](items []T, db gorm.DB) {
	result := db.Save(items)
	if result.Error != nil {

		typeName := reflect.TypeOf(items).Elem().Name()
		fmt.Println("failed to Batch Save", typeName)
		panic(result.Error)
	}
}

func BatchUpdateField[T mo.TodoPointers](items []T, field string, db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		insertDB := fs.Reduce3Args[T](UpdateFieldTx, tx)
		insertField := fs.Reduce2Args[T](insertDB, field)
		errors := fs.Map(items, insertField)
		anyError := fs.Reduce(errors, fs.CheckErr)
		return anyError
	})
	if err != nil {
		typeName := reflect.TypeOf(items).Elem().Name()
		fmt.Println("failed to batch update", typeName)
		panic(err)
	}
}

func UpdateField[T mo.TodoPointers](item T, field string, db *gorm.DB) {
	err := UpdateFieldTx(item, field, db)
	if err != nil {
		typeName := reflect.TypeOf(item).Elem().Name()
		fmt.Printf("failed to update field:[%s] for: [%s]", field, typeName)
		panic(err)
	}
}

func Save[T mo.TodoPointers](item T, db *gorm.DB) {
	result := db.Save(item)
	if result.Error != nil {
		typeName := reflect.TypeOf(item).Elem().Name()
		fmt.Println("failed to save", typeName)
		panic(result.Error)
	}
}
