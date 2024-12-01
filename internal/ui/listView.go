package ui

import (
	"fmt"

	mo "todo-cli/internal/models"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) renderListView() string {
	colorizer := mo.NewColorizer()
	s := colorizer.Colors["purple"] + colorizer.FontStyles["bold"] + "Todo List" + colorizer.Commands["reset"] + "\n\n"

	for i, choice := range m.Todos {

		title := choice.Title
		Cursor := " " // no cursor
		if m.Cursor == i {
			Cursor = colorizer.Colors["pink"] + colorizer.FontStyles["bold"] + ">" + colorizer.Commands["reset"] // cursor!
			title = colorizer.Colors["thickGreen"] + colorizer.FontStyles["bold"] + choice.Title + colorizer.Commands["reset"]
		}

		checked := " " // not selected
		if choice.Done {
			checked = "x" // selected!
		}

		s += fmt.Sprintf("%s [%s] %s\n", Cursor, checked, title)
	}

	s += "\n| q: quit | d: delete | o: add |\n"

	return s
}

func (m Model) updateListView(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q": // quit
			cmd = tea.Quit

		case "up", "k": // move cursor up
			if m.Cursor > 0 {
				m.Cursor--
			} else {
				m.Cursor = len(m.Todos) - 1
			}

		case "down", "j": // move cursor down
			if m.Cursor < len(m.Todos)-1 {
				m.Cursor++
			} else {
				m.Cursor = 0
			}

		case "enter", " ": // toggle done
			ok := m.Todos[m.Cursor].Done
			if ok {
				m.Todos[m.Cursor].Done = false
			} else {
				m.Todos[m.Cursor].Done = true
			}
			m.Repo.Update(&m.Todos[m.Cursor])

		case "o": // add new item
			newTodo := mo.Todo{Done: false, Title: "", Index: m.Cursor + 1, ID: 0}
			insertAtIndex(&m.Todos, m.Cursor+1, newTodo)
			m.Cursor++
			m.ActiveView = mo.Add
			if len(m.Todos) == 1 {
				m.Cursor = 0
			}

		case "d": // delete item
			m.Repo.Delete(int(m.Todos[m.Cursor].ID))
			m.Todos = append(m.Todos[:m.Cursor], m.Todos[m.Cursor+1:]...)
			if m.Cursor > len(m.Todos)-1 {
				m.Cursor = len(m.Todos) - 1
			}

		}
	}
	return m, cmd
}

func insertAtIndex[Type any](lisa *[]Type, index int, value Type) {
	*lisa = append(*lisa, value)
	if index >= len(*lisa) || index < 0 || len(*lisa) <= 1 {
		return
	}
	for i := len(*lisa) - 1; i > index; i-- {
		(*lisa)[i], (*lisa)[i-1] = (*lisa)[i-1], (*lisa)[i]
	}
}
