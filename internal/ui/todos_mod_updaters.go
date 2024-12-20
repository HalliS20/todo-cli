package ui

import (
	"sort"
	mo "todo-cli/internal/models"

	"github.com/charmbracelet/bubbletea"
)

func fixIndexing(lisa *[]mo.Todo) {
	for i := 0; i < len(*lisa); i++ {
		(*lisa)[i].Index = i
	}
}

func (m *Model) updateAddView(msg tea.Msg) (tea.Model, tea.Cmd) {
	todo := m.Todos[m.Cursor]
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg: // detects key press
		switch msg.String() {

		case "ctrl+c":
			m.ActiveView = mo.List
			m.Todos = append(m.Todos[:m.Cursor], m.Todos[m.Cursor+1:]...)
			cmd = tea.Quit

		case "enter":
			fixIndexing(&m.Todos)
			m.Repo.Todos.OrderAndAdd(&m.Todos)
			sort.Sort(mo.ByIndex(m.Todos))
			m.ActiveView = mo.List

		case "esc":
			m.ActiveView = mo.List
			m.Todos = append(m.Todos[:m.Cursor], m.Todos[m.Cursor+1:]...)
			m.Cursor--

		default:
			m.HandleTodoInput(msg, &todo)
			m.Todos[m.Cursor] = todo
		}
	}
	return m, cmd
}

func (m *Model) updateEditView(msg tea.Msg) (tea.Model, tea.Cmd) {
	todo := m.Todos[m.Cursor]
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg: // detects key press
		switch msg.String() {

		case "ctrl+c":
			m.ActiveView = mo.List
			cmd = tea.Quit

		case "enter":
			m.Repo.Todos.UpdateField(&todo, "Title")
			m.ActiveView = mo.List

		case "esc":
			m.ActiveView = mo.List
			m.Todos[m.Cursor].Title = m.EditCache

		default:
			m.HandleTodoInput(msg, &todo)
			m.Todos[m.Cursor] = todo

		}
	}
	return m, cmd
}
