package ui

import (
	op "todo-cli/internal/enums/operation"
	td "todo-cli/internal/models/todo"
	f "todo-cli/pkg/functions"

	tea "github.com/charmbracelet/bubbletea"
)

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

func (m *Model) HandleModification(msg tea.KeyMsg) {
	switch msg.String() {
	case "enter", " ": // toggle done
		if !m.ShownTodos[m.Cursor].Dir {
			m.ShownTodos[m.Cursor].Done = !m.ShownTodos[m.Cursor].Done
			m.Repo.Todos.UpdateField(m.ShownTodos[m.Cursor], "Done")
		} else {
			idCopy := m.ShownTodos[m.Cursor].ID
			m.ParentID = &idCopy
			nameCopy := m.ShownTodos[m.Cursor].Title
			m.ParentName = &nameCopy
			m.SwitchList()
		}

	case "o": // add new item
		m.Cursor++
		newTodo := td.Todo{Done: false, Title: "", Index: m.Cursor, ParentID: *m.ParentID}
		f.InsertAtIndex(&m.ShownTodos, m.Cursor, &newTodo)
		m.AllTodos = append(m.AllTodos, &newTodo)
		m.ActiveOp = op.Add
		if len(m.ShownTodos) == 1 {
			m.Cursor = 0
		}

	case "i": // edit item
		if len(m.ShownTodos) == 0 {
			return
		}
		m.ActiveOp = op.Edit
		m.EditCache = m.ShownTodos[m.Cursor].Title

	case "d", "backspace": // delete item
		m.BigRem(m.ShownTodos[m.Cursor])
		if m.ShownTodos[m.Cursor].Dir {
			m.DirDelete(m.ShownTodos[m.Cursor].ID)
		}
		m.Repo.Todos.Delete(int(m.ShownTodos[m.Cursor].ID))
		m.ShownTodos = append(m.ShownTodos[:m.Cursor], m.ShownTodos[m.Cursor+1:]...)
		if m.Cursor > len(m.ShownTodos)-1 {
			m.Cursor = len(m.ShownTodos) - 1
		}
	}
}

func (m *Model) HandleOrdering(msg tea.KeyMsg) {
	if len(m.ShownTodos) <= 1 {
		return
	}
	switch msg.String() {
	case "ctrl+k": // Operation: move item up
		// Visual: item moves up index goes down
		if m.Cursor > 0 {
			m.Repo.Todos.SwapListItems(&m.ShownTodos, m.Cursor, m.Cursor-1)
			m.Cursor--
		} else {
			// Exception:
			// cursor is at the list top = lowest index
			// we wrap the item around = bottom
			ListWrap(&m.ShownTodos, true)
			m.Repo.Todos.BatchUpdateField(&m.ShownTodos, "Index")
			m.Cursor = len(m.ShownTodos) - 1
			return
		}

	case "ctrl+j": // Operation:  move item down
		// Visual: item moves down index goes up
		if m.Cursor < len(m.ShownTodos)-1 {
			m.Repo.Todos.SwapListItems(&m.ShownTodos, m.Cursor, m.Cursor+1)
			m.Cursor++
		} else {
			//// Exception:
			//// cursor is at the list Bottom = highest index
			//// we wrap the item around = top
			ListWrap(&m.ShownTodos, false)
			m.Repo.Todos.BatchUpdateField(&m.ShownTodos, "Index")
			m.Cursor = 0
			return
		}
	}
}

func (m *Model) HandleNavUp(msg tea.KeyMsg) {
	switch msg.String() {
	case "-":
		pp := m.GetParentsParent()
		pi := m.getItem(pp)
		piName := "Todo List"
		if pi != nil {
			piName = pi.Title
		}
		*m.ParentID = pp
		*m.ParentName = piName
		m.SwitchList()
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

func (m *Model) DirDelete(id uint) {
	for _, item := range m.AllTodos {
		if item.ParentID != id {
			continue
		}
		if item.Dir {
			m.DirDelete(item.ID)
		} else {
			m.Repo.Todos.Delete(int(item.ID))
		}
	}
}

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
