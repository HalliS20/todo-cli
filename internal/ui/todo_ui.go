package ui

import (
	"todo-cli/internal/interfaces"
	mo "todo-cli/internal/models"

	"github.com/charmbracelet/bubbletea"
)

type Model struct {
	Todos      []mo.Todo
	Cursor     int
	Repo       interfaces.IRepository
	ActiveView mo.ActiveView
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.ActiveView == mo.List {
		return m.updateListView(msg)
	} else if m.ActiveView == mo.Add {
		return m.updateAddView(msg)
	} else {
		return m.updateListView(msg)
	}
}

func (m Model) View() string {
	if m.ActiveView == mo.List {
		return m.renderListView()
	} else if m.ActiveView == mo.Add {
		return m.renderListView()
	} else {
		return m.renderListView()
	}
}
