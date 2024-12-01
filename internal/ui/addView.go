package ui

import (
	"sort"
	mo "todo-cli/internal/models"
	"unicode"

	"github.com/charmbracelet/bubbletea"
)

func (m Model) renderAddView() string {
	s := m.renderListView()
	title := m.Todos[m.Cursor].Title
	s += "\n"
	s += title
	return s
}

func (m Model) updateAddView(msg tea.Msg) (tea.Model, tea.Cmd) {
	todo := m.Todos[m.Cursor]
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg: // detects key press
		switch msg.String() {

		case "ctrl+c":
			m.ActiveView = mo.Empty
			m.Todos = m.Todos[:len(m.Todos)-1]
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

func fixIndexing(lisa *[]mo.Todo) {
	for i := 0; i < len(*lisa); i++ {
		(*lisa)[i].Index = i
	}
}
