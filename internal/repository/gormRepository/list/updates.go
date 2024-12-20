package listRepo

import (
	b "todo-cli/internal/repository/gormRepository/gormBase"
	f "todo-cli/pkg/functions"
)

func (r *GormListRepo) UpdateField(todo *TodoList, field string) {
	b.UpdateField(todo, field, r.db)
}

func (r *GormListRepo) BatchUpdateField(todos *[]TodoList, field string) {
	pointerList := f.LPtoPL(todos)
	b.BatchUpdateField(pointerList, field, r.db)
}
