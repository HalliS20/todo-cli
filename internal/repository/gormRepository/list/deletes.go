package listRepo

import (
	"fmt"
	mo "todo-cli/internal/models/lists"
	"todo-cli/internal/models/todo"

	"gorm.io/gorm"
)

func (r *GormListRepo) Delete(id int) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Delete all todos for this list
		if err := tx.Where("list_id = ?", id).Delete(&todo.Todo{}).Error; err != nil {
			fmt.Println("failed to delete todos for list", id)
			return err
		}

		// Delete the list itself
		var zero mo.TodoList
		if err := tx.Delete(&zero, id).Error; err != nil {
			fmt.Println("failed to delete list", id)
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println("transaction failed:", err)
		panic(err)
	}
}
