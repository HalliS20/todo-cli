package repository

import (
	"fmt"
	"reflect"
	"strings"
	"todo-cli/internal/interfaces"
	"todo-cli/internal/models"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

type Todo = models.Todo

type GormDbRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) interfaces.IRepository {
	if db.AutoMigrate(&Todo{}) != nil {
		panic("failed to auto migrate")
	}
	return &GormDbRepository{db}
}

func (r *GormDbRepository) Create(todo *Todo) {
	result := r.db.Create(todo)
	if result.Error != nil {
		fmt.Println("failed to create")
		panic(result.Error)
	}
}

func (r *GormDbRepository) Update(todo *Todo) {
	result := r.db.Save(&todo)
	if result.Error != nil {
		fmt.Println("failed to update")
		panic(result.Error)
	}
}

func (r *GormDbRepository) UpdateField(todo Todo, field string) error {
	return r.UpdateFieldTx(r.db, todo, field)
}

func (r *GormDbRepository) UpdateFieldTx(db *gorm.DB, todo Todo, field string) error {
	fieldVal := reflect.ValueOf(todo).FieldByName(field)
	if !fieldVal.IsValid() {
		return fmt.Errorf("field %s does not exist", field)
	}

	return db.Model(&Todo{}).
		Where("id = ?", todo.ID).
		Update(strings.ToLower(field), fieldVal.Interface()).
		Error
}

func (r *GormDbRepository) BatchUpdate(todos []Todo) {
	result := r.db.Save(&todos)
	if result.Error != nil {
		fmt.Println("failed to update")
		panic(result.Error)
	}
}

func (r *GormDbRepository) FixAndAdd(todos []Todo) {
	error := r.db.Transaction(func(tx *gorm.DB) error {
		for i, todo := range todos {
			todo.Index = i
			if todo.ID == 0 {
				err := tx.Create(&todo).Error
				if err != nil {
					fmt.Println("failed to Create")
					panic(err)
				}
				continue
			}
			err := r.UpdateFieldTx(tx, todo, "Index")
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

func (r *GormDbRepository) Delete(id int) {
	result := r.db.Delete(&Todo{}, id)
	if result.Error != nil {
		fmt.Println("failed to delete")
		panic(result.Error)
	}
}

func (r *GormDbRepository) GetAll() []Todo {
	var todos []Todo
	result := r.db.Find(&todos)
	if result.Error != nil {
		fmt.Println("failed to get items")
		panic(result.Error)
	}
	return todos
}

func (r *GormDbRepository) GetOne(id int) Todo {
	var todo Todo
	result := r.db.Where("id = ?", id).First(&todo)
	if result.Error != nil {
		fmt.Println("failed to get items")
		panic(result.Error)
	}
	return todo
}
