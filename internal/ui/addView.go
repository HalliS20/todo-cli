package ui

import (
	"sort"
	mo "todo-cli/internal/models"
	"unicode"

	"github.com/charmbracelet/bubbletea"
)

func fixIndexing(lisa *[]mo.Todo) {
	for i := 0; i < len(*lisa); i++ {
		(*lisa)[i].Index = i
	}
}

func (m *Model) handleTextInput(msg tea.KeyMsg, todo *mo.Todo) {
	switch msg.String() {
	case "backspace", "delete":
		if len(todo.Title) > 0 {
			todo.Title = todo.Title[:len(todo.Title)-1]
		}
	case "space", " ":
		todo.Title += " "
	default:
		if unicode.IsPrint(rune(msg.String()[0])) {
			todo.Title += msg.String()
		}
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
			m.Repo.FixAndAdd(m.Todos)
			todos := m.Repo.GetAll()
			sort.Sort(mo.ByIndex(todos))
			m.Todos = todos
			m.ActiveView = mo.List

		case "esc":
			m.ActiveView = mo.List
			m.Todos = append(m.Todos[:m.Cursor], m.Todos[m.Cursor+1:]...)
			m.Cursor--

		default:
			m.handleTextInput(msg, &todo)
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
			m.Repo.UpdateField(todo, "Title")
			m.ActiveView = mo.List

		case "esc":
			m.ActiveView = mo.List
			m.Todos[m.Cursor].Title = m.EditCache

		default:
			m.handleTextInput(msg, &todo)
			m.Todos[m.Cursor] = todo

		}
	}
	return m, cmd
}
