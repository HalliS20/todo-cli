package listRepo

import (
	b "todo-cli/internal/repository/gormRepository/gormBase"
)

func (r *GormListRepo) UpdateField(todo *TodoList, field string) {
	b.UpdateField(todo, field, r.db)
}

func (r *GormListRepo) BatchUpdateField(todos *[]*TodoList, field string) {
	b.BatchUpdateField(*todos, field, r.db)
}
