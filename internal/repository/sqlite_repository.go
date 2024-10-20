package repository

import (
	"fmt"
	"todo-cli/internal/interfaces"
	"todo-cli/internal/models"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

type Todo = models.Todo

type SQLiteRepository struct {
	db *gorm.DB
}

func NewSQLiteRepository(db *gorm.DB) interfaces.IRepository {
	if db.AutoMigrate(&Todo{}) != nil {
		panic("failed to auto migrate")
	}
	return &SQLiteRepository{db}
}

func (r *SQLiteRepository) Create(todo *Todo) {
	result := r.db.Create(todo)
	if result.Error != nil {
		fmt.Println("failed to create")
		panic(result.Error)
	}
}

func (r *SQLiteRepository) Update(todo *Todo) {
	result := r.db.Save(&todo)
	if result.Error != nil {
		fmt.Println("failed to update")
		panic(result.Error)
	}
}

func (r *SQLiteRepository) Delete(id int) {
	result := r.db.Delete(&Todo{}, id)
	if result.Error != nil {
		fmt.Println("failed to delete")
		panic(result.Error)
	}
}

func (r *SQLiteRepository) GetAll() []Todo {
	var todos []Todo
	result := r.db.Find(&todos)
	if result.Error != nil {
		fmt.Println("failed to get items")
		panic(result.Error)
	}
	return todos
}

func (r *SQLiteRepository) GetOne(id int) Todo {
	var todo Todo
	result := r.db.Where("id = ?", id).First(&todo)
	if result.Error != nil {
		fmt.Println("failed to get items")
		panic(result.Error)
	}
	return todo
}
