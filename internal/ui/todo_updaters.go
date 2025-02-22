package ui

import (
	"sort"
	op "todo-cli/internal/enums/operation"
	td "todo-cli/internal/models/todo"
	bh "todo-cli/internal/ui/base"

	"github.com/charmbracelet/bubbletea"
)

func (m *Model) updateAddView(msg tea.Msg) tea.Cmd {
	todo := m.ShownTodos[m.Cursor]
	switch msg := msg.(type) {
	case tea.KeyMsg: // detects key press
		switch msg.String() {

		case "ctrl+c":
			return tea.Quit

		case "enter":
			lasti := len(todo.Title) - 1
			lastLetter := string(todo.Title[lasti])
			if lastLetter == "/" {
				todo.Dir = true
				todo.Title = todo.Title[:lasti]
			}
			m.Repo.Todos.OrderAndAdd(&m.ShownTodos)
			sort.Sort(td.Todos(m.AllTodos))
			m.ActiveOp = op.Lister
		case "esc":
			m.ActiveOp = op.Lister
			m.ShownTodos[m.Cursor].ID = 0
			m.ShownTodos = append(m.ShownTodos[:m.Cursor], m.ShownTodos[m.Cursor+1:]...)
			m.DelUnfinished()
			m.Cursor--

		default:
			bh.HandleTextInput(msg, &todo.Title)
			m.ShownTodos[m.Cursor] = todo
		}
	}

	return nil
}

func (m *Model) updateEditView(msg tea.Msg) tea.Cmd {
	todo := m.ShownTodos[m.Cursor]
	switch msg := msg.(type) {
	case tea.KeyMsg: // detects key press
		switch msg.String() {

		case "ctrl+c":
			return tea.Quit

		case "enter":
			lasti := len(todo.Title) - 1
			lastLetter := string(todo.Title[lasti])
			if lastLetter == "/" && !todo.Dir {
				todo.Dir = true
				todo.Title = todo.Title[:lasti]
				m.Repo.Todos.UpdateField(todo, "Dir")
			} else {
				m.Repo.Todos.UpdateField(todo, "Title")
				m.ActiveOp = op.Lister
			}

		case "esc":
			m.ActiveOp = op.Lister
			m.ShownTodos[m.Cursor].Title = m.EditCache

		default:
			bh.HandleTextInput(msg, &todo.Title)
			m.ShownTodos[m.Cursor] = todo

		}
	}
	return nil
}

func (m *Model) updateListView(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		default:
			cmd = bh.HandleNavigation(msg, &m.Cursor, len(m.ShownTodos))
			if cmd == nil {
				m.HandleNavUp(msg)
			}
			m.HandleModification(msg)
			m.HandleOrdering(msg)
		}
	}
	return cmd
}
