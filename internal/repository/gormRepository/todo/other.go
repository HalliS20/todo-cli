package todoRepo

import (
	b "todo-cli/internal/repository/gormRepository/gormBase"
	f "todo-cli/pkg/functions"
)

func (r *GormTodoRepo) Delete(id int) {
	b.Delete[Todo](id, r.db)
}

func (r *GormTodoRepo) GetAll() []Todo {
	return b.GetAll[Todo](r.db)
}

func (r *GormTodoRepo) GetOne(id int) Todo {
	return b.GetOne[Todo](id, r.db)
}

func (r *GormTodoRepo) Create(todo *Todo) {
	b.Create(todo, r.db)
}

func (r *GormTodoRepo) OrderAndAdd(items *[]Todo) {
	for i := range *items {
		(*items)[i].Index = i
	}
	newList := f.LPtoPL(items)
	b.UpdateOrAdd(newList, r.db)
}
