package ui

import (
	mo "todo-cli/internal/models"
	f "todo-cli/pkg/functions"
	// "unicode"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) HandleTodoInput(msg tea.KeyMsg, item *mo.Todo) {
	handleTextInput(msg, &item.Title)
}

func (m *Model) HandleNavigation(msg tea.KeyMsg) tea.Cmd {
	return handleNavigation(msg, m, m.Todos)
}

func (m *Model) HandleModification(msg tea.KeyMsg) {
	switch msg.String() {
	case "enter", " ": // toggle done
		ok := m.Todos[m.Cursor].Done
		if ok {
			m.Todos[m.Cursor].Done = false
		} else {
			m.Todos[m.Cursor].Done = true
		}
		m.Repo.Todos.UpdateField(&m.Todos[m.Cursor], "Done")

	case "o": // add new item
		newTodo := mo.Todo{Done: false, Title: "", Index: m.Cursor + 1, ID: 0}
		f.InsertAtIndex(&m.Todos, m.Cursor+1, newTodo)
		m.Cursor++
		m.ActiveView = mo.Add
		if len(m.Todos) == 1 {
			m.Cursor = 0
		}

	case "i": // edit item
		m.ActiveView = mo.Edit
		m.EditCache = m.Todos[m.Cursor].Title

	case "d": // delete item
		m.Repo.Todos.Delete(int(m.Todos[m.Cursor].ID))
		m.Todos = append(m.Todos[:m.Cursor], m.Todos[m.Cursor+1:]...)
		if m.Cursor > len(m.Todos)-1 {
			m.Cursor = len(m.Todos) - 1
		}
	}
}

func (m *Model) HandleOrdering(msg tea.KeyMsg) {
	switch msg.String() {
	case "ctrl+k": // move item up
		if m.Cursor > 0 {
			m.Todos[m.Cursor].Index = m.Cursor - 1
			m.Todos[m.Cursor-1].Index = m.Cursor
			m.Todos[m.Cursor], m.Todos[m.Cursor-1] = m.Todos[m.Cursor-1], m.Todos[m.Cursor]
			modifiedTodos := []mo.Todo{m.Todos[m.Cursor], m.Todos[m.Cursor-1]}
			m.Repo.Todos.BatchUpdateField(&modifiedTodos, "Index")
			m.Cursor--
		} else {
			cache := m.Todos[m.Cursor]
			m.Cursor = len(m.Todos) - 1
			for i := 1; i < len(m.Todos); i++ {
				m.Todos[i].Index = i - 1
				m.Todos[i-1] = m.Todos[i]
			}
			m.Todos[len(m.Todos)-1] = cache
			m.Todos[len(m.Todos)-1].Index = len(m.Todos) - 1
			m.Repo.Todos.BatchUpdateField(&m.Todos, "Index")
		}

	case "ctrl+j": // move item down
		if m.Cursor < len(m.Todos)-1 {
			m.Todos[m.Cursor+1].Index = m.Cursor
			m.Todos[m.Cursor].Index = m.Cursor + 1
			m.Todos[m.Cursor], m.Todos[m.Cursor+1] = m.Todos[m.Cursor+1], m.Todos[m.Cursor]
			modifiedTodos := []mo.Todo{m.Todos[m.Cursor], m.Todos[m.Cursor+1]}
			m.Repo.Todos.BatchUpdateField(&modifiedTodos, "Index")
			m.Cursor++
		} else {
			cache := m.Todos[m.Cursor]
			m.Cursor = 0
			for i := len(m.Todos) - 1; i > 0; i-- {
				m.Todos[i].Index = i + 1
				m.Todos[i] = m.Todos[i-1]
			}
			m.Todos[0] = cache
			m.Todos[0].Index = 0
			m.Repo.Todos.BatchUpdateField(&m.Todos, "Index")
		}
	}
}
