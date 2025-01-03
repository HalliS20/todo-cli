package todoRepo

import (
	b "todo-cli/internal/repository/gormRepository/gormBase"
)

func (r *GormTodoRepo) UpdateField(todo *Todo, field string) {
	b.UpdateField(todo, field, r.db)
}

func (r *GormTodoRepo) BatchUpdateField(todos *[]*Todo, field string) {
	b.BatchUpdateField(*todos, field, r.db)
}
