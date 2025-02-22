package ui

import (
	"strconv"
	td "todo-cli/internal/models/todo"
)

func ListWrap(lisa *[]*td.Todo, up bool) {
	// if we go up we are moving from the end(0) to one from back (listEnd - 1)
	listEnd := len(*lisa) - 1

	end := 0
	oneFromOther := listEnd - 1

	// \/\/\/\/ Swap front and back \/\/\/\/ \\
	(*lisa)[0], (*lisa)[listEnd] = (*lisa)[listEnd], (*lisa)[0]
	(*lisa)[0].Index, (*lisa)[listEnd].Index = (*lisa)[listEnd].Index, (*lisa)[0].Index
	// now we have the one that we want to move at the other end
	// but the collateral item is at the other end when it should be next to the moved item

	if listEnd <= 1 {
		// if there are only 2 items we just need the swap
		return
	}
	if !up {
		// if we go down we are moving from the end(listEnd) to one from front (1)
		end = listEnd
		oneFromOther = 1
	}
	HoistItem(lisa, end, oneFromOther)
}

func HoistItem(lisa *[]*td.Todo, from int, to int) {
	n := 1
	if to < from {
		n = -1
	}
	end := len(*lisa)
	for i := from; i != to; i += n {
		if i+n < 0 || i+n >= end || i < 0 || i >= end {
			panic("hoisting went out of bounds at i: " + strconv.Itoa(i) + " n:" + strconv.Itoa(n))
		}
		(*lisa)[i], (*lisa)[i+n] = (*lisa)[i+n], (*lisa)[i]
		(*lisa)[i].Index, (*lisa)[i+n].Index = (*lisa)[i+n].Index, (*lisa)[i].Index
	}
}

func (m *Model) getItem(id uint) *td.Todo {
	for _, item := range m.AllTodos {
		if item.ID == id {
			return item
		}
	}
	return nil
}

func (m *Model) GetParentsParent() uint {
	for _, item := range m.AllTodos {
		if item.ID == *m.ParentID {
			return item.ParentID
		}
	}
	return 0
}

func (m *Model) SwitchList() {
	m.ShownTodos = []*td.Todo{}
	for _, item := range m.AllTodos {
		if item.ParentID == *m.ParentID {
			m.ShownTodos = append(m.ShownTodos, item)
		}
	}
	if m.Cursor >= len(m.ShownTodos) {
		m.Cursor = 0
	}
}

// ========== THESE 2 Delete functions may be doing the same thing
// ========== too lazy to find out right now, its late
func (m *Model) DirDelete(id uint) {
	for _, item := range m.AllTodos {
		if item.ParentID != id {
			continue
		}
		if item.Dir {
			m.DirDelete(item.ID)
			m.Repo.Todos.Delete(int(item.ID))
		} else {
			m.Repo.Todos.Delete(int(item.ID))
		}
	}
}

func (m *Model) BigRem(todo *td.Todo) {
	shiftMode := false
	i := todo.Index
	for ; i < len(m.AllTodos)-1; i++ {
		if todo.ID == m.AllTodos[i].ID {
			shiftMode = true
		}
		if shiftMode {
			m.AllTodos[i] = m.AllTodos[i+1]
		}
	}
	if !shiftMode && i != len(m.AllTodos)-1 {
		panic("id not found")
	}
	m.AllTodos = m.AllTodos[:len(m.AllTodos)-1]
}

//====================== :)

func (m *Model) DelUnfinished() {
	i := 0
	item := &td.Todo{}
	for i, item = range m.AllTodos {
		if item.ID == 0 {
			break
		}
	}
	lennie := len(m.AllTodos)
	if i != len(m.AllTodos)-1 {
		m.AllTodos = append(m.AllTodos[:i], m.AllTodos[i+1:]...)
		return
	}
	m.AllTodos = m.AllTodos[:lennie-1]
}
