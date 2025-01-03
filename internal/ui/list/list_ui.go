package list_ui

import (
	"sort"
	cm "todo-cli/internal/enums/command"
	op "todo-cli/internal/enums/operation"
	ls "todo-cli/internal/models/lists"
	vw "todo-cli/internal/models/view"
	sqli "todo-cli/internal/repository/gormRepository"
	cl "todo-cli/pkg/colorizer"

	"github.com/charmbracelet/bubbletea"
)

type Model struct {
	ActiveOp  op.ActiveOp
	Repo      sqli.GormRepository
	Lists     []*ls.TodoList
	ListID    *uint
	ListName  *string
	Cursor    int
	Views     map[op.ActiveOp]vw.View
	EditCache string
}

func NewListModel(repo sqli.GormRepository) *Model {
	lisa := repo.Lists.GetAll()
	sort.Sort(ls.Lists(lisa))
	initialID := uint(0)
	initialName := ""
	newMdl := Model{
		Repo:     repo,
		ActiveOp: op.Lister,
		Views:    make(map[op.ActiveOp]vw.View),
		Lists:    lisa,
		ListID:   &initialID,
		ListName: &initialName,
		Cursor:   0,
	}
	newMdl.Init()
	return &newMdl
}

func (m *Model) Init() tea.Cmd {
	m.Views[op.Lister] = vw.View{
		Update:  m.updateListView,
		View:    m.renderListView,
		Footer:  "\n| q: quit | d: delete | o: add | i: edit |\n",
		Header:  "Normal",
		OpColor: cl.White,
	}

	m.Views[op.Add] = vw.View{
		Update:  m.updateAddView,
		View:    m.renderListView,
		Footer:  "\n| ctrl+c: quit | enter: save | esc: back |",
		Header:  "Add",
		OpColor: cl.Green,
	}

	m.Views[op.Edit] = vw.View{
		Update:  m.updateEditView,
		View:    m.renderListView,
		Footer:  "\n| ctrl+c: quit | enter: save | esc: back |",
		Header:  "Edit",
		OpColor: cl.Yellow,
	}

	return nil
}

func (m *Model) Update(msg tea.Msg) (*Model, cm.Command) {
	return m, m.Views[m.ActiveOp].Update(msg)
}

func (m *Model) View() string {
	ss := m.Views[m.ActiveOp].View()
	ss += m.Views[m.ActiveOp].Footer
	return ss
}
