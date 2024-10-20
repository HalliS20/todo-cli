package main

import (
	"fmt"
	"os"
	"todo-cli/internal/interfaces"
	"todo-cli/internal/models"
	"todo-cli/internal/repository"
	"todo-cli/pkg"

	"github.com/charmbracelet/bubbletea"
)

type Todo = models.Todo

type model struct {
	todos  []Todo
	cursor int
	repo   interfaces.IRepository
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			ok := m.todos[m.cursor].Done
			if ok {
				m.todos[m.cursor].Done = false
			} else {
				m.todos[m.cursor].Done = true
			}
			m.repo.Update(&m.todos[m.cursor])

		// The "d" key deletes the item that the cursor is pointing at.
		case "d":
			m.repo.Delete(int(m.todos[m.cursor].ID))
			m.todos = append(m.todos[:m.cursor], m.todos[m.cursor+1:]...)
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Todos\n\n"

	// Iterate over our choices
	for i, choice := range m.todos {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if choice.Done {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Title)
	}

	// The footer
	s += "\n| q: quit | d: delete |\n"

	// Send the UI for rendering
	return s
}

func initialModel() tea.Model {
	db, err := pkg.OpenSqLiteDatabase("./db/todo.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	repo := repository.NewSQLiteRepository(db)
	todoList := repo.GetAll()

	return model{
		todos: todoList,
		repo:  repo,
	}
}

func main() {
	initModel := initialModel()
	p := tea.NewProgram(initModel)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
