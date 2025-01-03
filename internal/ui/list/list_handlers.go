package list_ui

import (
	cm "todo-cli/internal/enums/command"
	op "todo-cli/internal/enums/operation"
	mo "todo-cli/internal/models/lists"
	f "todo-cli/pkg/functions"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) HandleModification(msg tea.KeyMsg) {
	switch msg.String() {

	case "o": // add new item
		newList := mo.TodoList{Title: "", Index: m.Cursor + 1, ID: 0}
		f.InsertAtIndex(&m.Lists, m.Cursor+1, &newList)
		m.Cursor++
		m.ActiveOp = op.Add
		if len(m.Lists) == 1 {
			m.Cursor = 0
		}

	case "i": // edit item
		if len(m.Lists) == 0 {
			return
		}
		m.ActiveOp = op.Edit
		m.EditCache = m.Lists[m.Cursor].Title

	case "d", "backspace": // delete item
		m.Repo.Lists.Delete(int(m.Lists[m.Cursor].ID))
		m.Lists = append(m.Lists[:m.Cursor], m.Lists[m.Cursor+1:]...)
		if m.Cursor > len(m.Lists)-1 {
			m.Cursor = len(m.Lists) - 1
		}
	}
}

func (m *Model) HandleOrdering(msg tea.KeyMsg) {
	if len(m.Lists) <= 1 {
		return
	}

	switch msg.String() {
	case "ctrl+k": // move item up
		if m.Cursor > 0 {
			m.Repo.Lists.SwapListItems(&m.Lists, m.Cursor, m.Cursor-1)
			m.Cursor--
		} else {
			// For a description of this wrapping problem check the todo_ui.go file
			ListWrap(&m.Lists, true)
			m.Repo.Lists.BatchUpdateField(&m.Lists, "Index")
			m.Cursor = len(m.Lists) - 1

		}

	case "ctrl+j": // move item down
		if m.Cursor < len(m.Lists)-1 {
			m.Repo.Lists.SwapListItems(&m.Lists, m.Cursor, m.Cursor+1)
			m.Cursor++
		} else {
			// For a description of this wrapping problem check the todo_ui.go file
			ListWrap(&m.Lists, false)
			m.Repo.Lists.BatchUpdateField(&m.Lists, "Index")
			m.Cursor = 0

		}
	}
}

func (m *Model) HandleEnteringList(msg tea.KeyMsg) cm.Command {
	switch msg.String() {
	case "enter":
		*m.ListID = m.Lists[m.Cursor].ID
		*m.ListName = m.Lists[m.Cursor].Title
		return cm.NavTo
	}
	return cm.None
}
