package ui

import (
	mo "todo-cli/internal/models"
	f "todo-cli/pkg/functions"
	"unicode"

	"github.com/charmbracelet/bubbletea"
)

func handleTextInput(msg tea.KeyMsg, textField *string) {
	switch msg.String() {
	case "backspace", "delete":
		if len(*textField) > 0 {
			*textField = (*textField)[:len(*textField)-1]
		}
	case "space", " ":
		*textField += " "
	default:
		if unicode.IsPrint(rune(msg.String()[0])) {
			*textField += msg.String()
		}
	}
}

func handleNavigation[T mo.TodoTypes](msg tea.KeyMsg, mod *Model, navList []T) tea.Cmd {
	switch msg.String() {
	case "ctrl+c", "q": // quit
		return tea.Quit
	case "up", "k": // move cursor up
		if mod.Cursor > 0 {
			mod.Cursor--
		} else {
			mod.Cursor = len(navList) - 1
		}
	case "down", "j": // move cursor down
		if mod.Cursor < len(navList)-1 {
			mod.Cursor++
		} else {
			mod.Cursor = 0
		}
	}
	return nil
}

func HandleModification(msg tea.KeyMsg, m *Model) {
	switch msg.String() {
	case "enter", " ": // toggle done
		ok := m.Todos[m.Cursor].Done
		if ok {
			m.Todos[m.Cursor].Done = false
		} else {
			m.Todos[m.Cursor].Done = true
		}
		m.Repo.Todos.UpdateField(&m.Todos[m.Cursor], "Done")

	case "o": // add new item
		newTodo := mo.Todo{Title: "", Index: m.Cursor + 1, ID: 0}
		f.InsertAtIndex(&m.Todos, m.Cursor+1, newTodo)
		m.Cursor++
		m.ActiveView = mo.Add
		if len(m.Todos) == 1 {
			m.Cursor = 0
		}

	case "i": // edit item
		m.ActiveView = mo.Edit
		m.EditCache = m.Todos[m.Cursor].Title

	case "d": // delete item
		m.Repo.Todos.Delete(int(m.Todos[m.Cursor].ID))
		m.Todos = append(m.Todos[:m.Cursor], m.Todos[m.Cursor+1:]...)
		if m.Cursor > len(m.Todos)-1 {
			m.Cursor = len(m.Todos) - 1
		}
	}
}
