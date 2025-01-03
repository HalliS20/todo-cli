package todo_ui

import (
	"sort"
	cm "todo-cli/internal/enums/command"
	op "todo-cli/internal/enums/operation"
	td "todo-cli/internal/models/todo"
	vw "todo-cli/internal/models/view"
	sqli "todo-cli/internal/repository/gormRepository"
	cl "todo-cli/pkg/colorizer"

	"github.com/charmbracelet/bubbletea"
)

type Model struct {
	ActiveOp   op.ActiveOp
	Repo       sqli.GormRepository
	AllTodos   []*td.Todo
	ShownTodos []*td.Todo
	ListID     *uint
	ListName   *string
	Cursor     int
	Views      map[op.ActiveOp]vw.View
	EditCache  string
}

func NewTodoModel(repo sqli.GormRepository) *Model {
	lisa := repo.Todos.GetAll()
	sort.Sort(td.Todos(lisa))
	initialID := uint(0)
	initialName := ""
	newMdl := Model{
		Repo:       repo,
		AllTodos:   lisa,
		ShownTodos: []*td.Todo{},
		ActiveOp:   op.Lister,
		Views:      make(map[op.ActiveOp]vw.View),
		ListID:     &initialID,
		ListName:   &initialName,
	}
	newMdl.Init()
	return &newMdl
}

func (m *Model) Init() tea.Cmd {
	m.Views[op.Lister] = vw.View{
		Update:  m.updateListView,
		View:    m.renderListView,
		Footer:  "\n| q : quit | d : delete | o : add | i : edit | - : back |\n",
		Header:  "Normal",
		OpColor: cl.White,
	}

	m.Views[op.Add] = vw.View{
		Update:  m.updateAddView,
		View:    m.renderListView,
		Footer:  "\n| ctrl+c : quit | enter : save | esc : back |",
		Header:  "Add",
		OpColor: cl.Green,
	}

	m.Views[op.Edit] = vw.View{
		Update:  m.updateEditView,
		View:    m.renderListView,
		Footer:  "\n| ctrl+c : quit | enter : save | esc : back |",
		Header:  "Edit",
		OpColor: cl.Yellow,
	}

	m.SwitchList()

	return nil
}

func (m *Model) Update(msg tea.Msg) (*Model, cm.Command) {
	return m, m.Views[m.ActiveOp].Update(msg)
}

func (m Model) View() string {
	ss := m.Views[m.ActiveOp].View()
	ss += m.Views[m.ActiveOp].Footer
	return ss
}

func (m *Model) SwitchList() {
	m.ShownTodos = []*td.Todo{}
	for _, item := range m.AllTodos {
		if item.ListID == *m.ListID {
			m.ShownTodos = append(m.ShownTodos, item)
		}
	}
}
