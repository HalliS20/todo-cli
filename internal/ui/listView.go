package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	mo "todo-cli/internal/models"
)

func (m Model) renderListView() string {
	colorizer := mo.NewColorizer()
	s := colorizer.Colors["purple"] + colorizer.FontStyles["bold"] + "Todo List" + colorizer.Commands["reset"] + "\n\n"
	// s += "\nTodos\n\n"

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

	s += "\n| q: quit | d: delete | a: add |\n"

	return s
}

func (m Model) updateListView(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			cmd = tea.Quit

		// The "up" and "k" keys move the Cursor up
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			} else {
				m.Cursor = len(m.Todos) - 1
			}

		// The "down" and "j" keys move the Cursor down
		case "down", "j":
			if m.Cursor < len(m.Todos)-1 {
				m.Cursor++
			} else {
				m.Cursor = 0
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			ok := m.Todos[m.Cursor].Done
			if ok {
				m.Todos[m.Cursor].Done = false
			} else {
				m.Todos[m.Cursor].Done = true
			}
			m.Repo.Update(&m.Todos[m.Cursor])

		case "o":
			newTodo := mo.Todo{Done: false, Title: ""}
			m.Todos = append(m.Todos, newTodo)
			m.Cursor = len(m.Todos) - 1
			m.ActiveView = mo.Add

		// The "d" key deletes the item that the Cursor is pointing at.
		case "d":
			m.Repo.Delete(int(m.Todos[m.Cursor].ID))
			m.Todos = append(m.Todos[:m.Cursor], m.Todos[m.Cursor+1:]...)
			if m.Cursor > len(m.Todos)-1 {
				m.Cursor = len(m.Todos) - 1
			}

		case "a":
			m.ActiveView = mo.Add
			m.Todos = append(m.Todos, mo.Todo{Done: false, Title: ""})
			m.Cursor = len(m.Todos) - 1
		}
	}

	// Return the updated Model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, cmd
}
