package listRepo

import (
	b "todo-cli/internal/repository/gormRepository/gormBase"
	f "todo-cli/pkg/functions"
)

func (r *GormListRepo) Delete(id int) {
	b.Delete[TodoList](id, r.db)
}

func (r *GormListRepo) GetAll() []TodoList {
	return b.GetAll[TodoList](r.db)
}

func (r *GormListRepo) GetOne(id int) TodoList {
	return b.GetOne[TodoList](id, r.db)
}

func (r *GormListRepo) Create(todo *TodoList) {
	b.Create(todo, r.db)
}

func (r *GormListRepo) OrderAndAdd(items *[]TodoList) {
	for i := range *items {
		(*items)[i].Index = i
	}
	newList := f.LPtoPL(items)
	b.UpdateOrAdd(newList, r.db)
}
