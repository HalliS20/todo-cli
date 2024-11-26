package ui

import (
	mo "todo-cli/internal/models"
	"unicode"

	"github.com/charmbracelet/bubbletea"
)

func (m Model) renderAddView() string {
	todo := m.Todos[m.Cursor]
	s := "Add a new todo\n\n"
	s += "   " + "Title: " + todo.Title + "\n\n"

	s += "\n| ctrl+c: quit | enter: add | esc : cancel |\n"

	return s
}

func (m Model) updateAddView(msg tea.Msg) (tea.Model, tea.Cmd) {
	todo := m.Todos[m.Cursor]
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			m.ActiveView = mo.Empty
			m.Todos = m.Todos[:len(m.Todos)-1]
			cmd = tea.Quit

		case "enter":
			m.ActiveView = mo.List
			m.Repo.Create(&todo)
			m.Todos = m.Repo.GetAll()

		case "esc":
			m.ActiveView = mo.List
			m.Todos = m.Todos[:len(m.Todos)-1]
			m.Cursor--

		case "backspace", "delete":
			if len(todo.Title) > 0 {
				todo.Title = todo.Title[:len(todo.Title)-1]
				m.Todos[m.Cursor] = todo
			}

		case "space", " ":
			todo.Title += " "
			m.Todos[m.Cursor] = todo

		default:
			if unicode.IsPrint(rune(msg.String()[0])) {
				todo.Title += msg.String()
				m.Todos[m.Cursor] = todo
			}
		}
	}
	return m, cmd
}
