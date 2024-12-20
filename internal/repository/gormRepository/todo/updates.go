package todoRepo

import (
	b "todo-cli/internal/repository/gormRepository/gormBase"
	f "todo-cli/pkg/functions"
)

func (r *GormTodoRepo) UpdateField(todo *Todo, field string) {
	b.UpdateField(todo, field, r.db)
}

func (r *GormTodoRepo) BatchUpdateField(todos *[]Todo, field string) {
	pointerList := f.LPtoPL(todos)
	b.BatchUpdateField(pointerList, field, r.db)
}
