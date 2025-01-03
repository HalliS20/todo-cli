package todoRepo

import (
	mo "todo-cli/internal/models/todo"
	b "todo-cli/internal/repository/gormRepository/gormBase"
)

func (r *GormTodoRepo) Delete(id int) {
	b.Delete[Todo](id, r.db)
}

func (r *GormTodoRepo) GetAll() []*Todo {
	return b.GetAll[Todo](r.db)
}

func (r *GormTodoRepo) GetForList(listID uint) []Todo {
	var todos []Todo
	r.db.Where("list_id = ?", listID).Find(&todos)
	return todos
}

func (r *GormTodoRepo) GetOne(id int) Todo {
	return b.GetOne[Todo](id, r.db)
}

func (r *GormTodoRepo) Create(todo *Todo) {
	b.Create(todo, r.db)
}

func (r *GormTodoRepo) OrderAndAdd(items *[]*Todo) {
	for i := range *items {
		(*items)[i].Index = i
	}
	b.UpdateOrAdd(items, r.db)
}

func (r *GormTodoRepo) SwapListItems(lisa *[]*mo.Todo, i1 int, i2 int) {
	(*lisa)[i1].Index, (*lisa)[i2].Index = (*lisa)[i2].Index, (*lisa)[i1].Index
	*(*lisa)[i2], *(*lisa)[i1] = *(*lisa)[i1], *(*lisa)[i2]
	modified := []*mo.Todo{(*lisa)[i1], (*lisa)[i2]}
	r.BatchUpdateField(&modified, "Index")
}
