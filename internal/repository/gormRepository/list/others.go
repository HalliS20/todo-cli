package listRepo

import (
	mo "todo-cli/internal/models/lists"
	b "todo-cli/internal/repository/gormRepository/gormBase"
)

func (r *GormListRepo) GetAll() []*TodoList {
	return b.GetAll[TodoList](r.db)
}

func (r *GormListRepo) GetOne(id int) TodoList {
	return b.GetOne[TodoList](id, r.db)
}

func (r *GormListRepo) Create(todo *TodoList) {
	b.Create(todo, r.db)
}

func (r *GormListRepo) OrderAndAdd(items *[]*TodoList) {
	for i := range *items {
		(*items)[i].Index = i
	}
	b.UpdateOrAdd[*TodoList](items, r.db)
}

func (r *GormListRepo) SwapListItems(lisa *[]*mo.TodoList, i1 int, i2 int) {
	(*lisa)[i1].Index, (*lisa)[i2].Index = (*lisa)[i2].Index, (*lisa)[i1].Index
	*(*lisa)[i2], *(*lisa)[i1] = *(*lisa)[i1], *(*lisa)[i2]
	modified := []*mo.TodoList{(*lisa)[i1], (*lisa)[i2]}
	r.BatchUpdateField(&modified, "Index")
}
