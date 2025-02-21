package ui

import (
	"sort"
	op "todo-cli/internal/enums/operation"
	td "todo-cli/internal/models/todo"
	vw "todo-cli/internal/models/view"
	sqli "todo-cli/internal/repository/gormRepository"
	cl "todo-cli/pkg/colorizer"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ActiveOp   op.ActiveOp
	Repo       sqli.GormRepository
	AllTodos   []*td.Todo
	ShownTodos []*td.Todo
	ParentID   *uint
	ParentName *string
	Cursor     int
	Views      map[op.ActiveOp]vw.View
	EditCache  string
}

func NewUI(repo sqli.GormRepository) *Model {
	lisa := repo.Todos.GetAll()
	sort.Sort(td.Todos(lisa))
	initialID := uint(0)
	initialName := "Todo List"
	newMdl := Model{
		Repo:       repo,
		AllTodos:   lisa,
		ShownTodos: []*td.Todo{},
		ActiveOp:   op.Lister,
		Views:      make(map[op.ActiveOp]vw.View),
		ParentID:   &initialID,
		ParentName: &initialName,
	}
	newMdl.Init()
	return &newMdl
}

func (ui *Model) Init() tea.Cmd {
	ui.Views[op.Lister] = vw.View{
		Update:  ui.updateListView,
		View:    ui.renderListView,
		Footer:  "\n| q : quit | d : delete | o : add | i : edit | - : back |\n",
		Header:  "Normal",
		OpColor: cl.White,
	}

	ui.Views[op.Add] = vw.View{
		Update:  ui.updateAddView,
		View:    ui.renderListView,
		Footer:  "\n| ctrl+c : quit | enter : save | esc : back |",
		Header:  "Add",
		OpColor: cl.Green,
	}

	ui.Views[op.Edit] = vw.View{
		Update:  ui.updateEditView,
		View:    ui.renderListView,
		Footer:  "\n| ctrl+c : quit | enter : save | esc : back |",
		Header:  "Edit",
		OpColor: cl.Yellow,
	}

	ui.SwitchList()

	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	action := m.Views[m.ActiveOp].Update(msg)
	if action != nil {
		action = tea.Sequence(func() tea.Msg { return nil }, tea.Quit)
	}
	return m, action
}

func (ui *Model) View() string {
	ss := ui.Views[ui.ActiveOp].View()
	ss += ui.Views[ui.ActiveOp].Footer
	return ss
}
