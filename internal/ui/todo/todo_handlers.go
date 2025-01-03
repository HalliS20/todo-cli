package todo_ui

import (
	cm "todo-cli/internal/enums/command"
	op "todo-cli/internal/enums/operation"
	mo "todo-cli/internal/models/todo"
	f "todo-cli/pkg/functions"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) BigRem(todo *mo.Todo) {
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
		ok := m.ShownTodos[m.Cursor].Done
		if ok {
			m.ShownTodos[m.Cursor].Done = false
		} else {
			m.ShownTodos[m.Cursor].Done = true
		}
		m.Repo.Todos.UpdateField(m.ShownTodos[m.Cursor], "Done")

	case "o": // add new item
		m.Cursor++
		newTodo := mo.Todo{Done: false, Title: "", Index: m.Cursor, ListID: *m.ListID}
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

func (m *Model) HandleNavUp(msg tea.KeyMsg) cm.Command {
	switch msg.String() {
	case "-":
		*m.ListID = 0
		m.Cursor = 0
		m.SwitchList()
		return cm.NavUp
	}
	return cm.None
}
