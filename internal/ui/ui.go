package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	vw "todo-cli/internal/enums/active_view"
	cm "todo-cli/internal/enums/command"
	sqli "todo-cli/internal/repository/gormRepository"
	ls "todo-cli/internal/ui/list"
	td "todo-cli/internal/ui/todo"
)

type UI struct {
	TodoModel  *td.Model
	ListModel  *ls.Model
	ActiveView vw.ActiveView
	Repo       sqli.GormRepository
	ListID     uint
	ListName   string
}

func NewUI(repo sqli.GormRepository) *UI {
	newUI := UI{
		ActiveView: vw.Lists,
		Repo:       repo,
		TodoModel:  td.NewTodoModel(repo),
		ListModel:  ls.NewListModel(repo),
		ListID:     0,
	}
	newUI.TodoModel.ListID = &newUI.ListID
	newUI.ListModel.ListID = &newUI.ListID
	newUI.ListModel.ListName = &newUI.ListName
	newUI.TodoModel.ListName = &newUI.ListName
	return &newUI
}

func (ui *UI) Init() tea.Cmd {
	return nil
}

func (ui *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	cmd := ui.updater(msg)
	switch cmd {
	case cm.Quit:
		ui.ActiveView = vw.Quit
		return ui, tea.Sequence(func() tea.Msg { return nil },
			tea.Quit)
	case cm.NavUp:
		ui.ActiveView = vw.Lists
	case cm.NavTo:
		ui.TodoModel.SwitchList()
		ui.ActiveView = vw.Todos
	case cm.None:
		// Do nothing
	}
	return ui, nil
}

func (ui *UI) updater(msg tea.Msg) cm.Command {
	var cmd cm.Command
	switch ui.ActiveView {
	case vw.Todos:
		_, cmd = ui.TodoModel.Update(msg)
	case vw.Lists:
		_, cmd = ui.ListModel.Update(msg)
	case vw.Quit:
		_, cmd = "", cm.Quit
	}
	return cmd
}

func (ui *UI) View() string {
	var viewStr string
	switch ui.ActiveView {
	case vw.Todos:
		viewStr = ui.TodoModel.View()
	case vw.Lists:
		viewStr = ui.ListModel.View()
	case vw.Quit:
		return ""
	}
	return viewStr
}
