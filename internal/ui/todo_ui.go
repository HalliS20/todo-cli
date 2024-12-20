package ui

import (
	"sort"
	mo "todo-cli/internal/models"
	sqli "todo-cli/internal/repository/gormRepository"

	"github.com/charmbracelet/bubbletea"
)

type ViewType struct {
	Update func(msg tea.Msg) (tea.Model, tea.Cmd)
	View   func() string
	Footer string
}

type Model struct {
	Todos      []mo.Todo
	Cursor     int
	Repo       sqli.GormRepository
	ActiveView mo.ActiveView
	Views      map[mo.ActiveView]ViewType
	EditCache  string
}

func NewModel(repo sqli.GormRepository) Model {
	lisa := repo.Todos.GetAll()
	sort.Sort(mo.ByIndex(lisa))

	return Model{
		Repo:       repo,
		Todos:      lisa,
		ActiveView: mo.List,
		Views:      make(map[mo.ActiveView]ViewType),
	}
}

func (m Model) Init() tea.Cmd {
	m.Views[mo.List] = ViewType{
		Update: m.updateListView,
		View:   m.renderListView,
		Footer: "\n| q: quit | d: delete | o: add | i: edit |\n",
	}

	m.Views[mo.Add] = ViewType{
		Update: m.updateAddView,
		View:   m.renderListView,
		Footer: "\n| ctrl+c: quit | enter: save | esc: back |",
	}

	m.Views[mo.Edit] = ViewType{
		Update: m.updateEditView,
		View:   m.renderListView,
		Footer: "\n| ctrl+c: quit | enter: save | esc: back |",
	}

	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.Views[m.ActiveView].Update(msg)
}

func (m Model) View() string {
	return m.Views[m.ActiveView].View() + m.Views[m.ActiveView].Footer
}
