package main

import (
	"fmt"
	"os"
	"todo-cli/internal/repository"
	"todo-cli/internal/ui"
	"todo-cli/pkg"

	"github.com/charmbracelet/bubbletea"
)

func initialModel() tea.Model {
	db, err := pkg.OpenSqLiteDatabase("~/.todo/db/todo.db", false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	repo := repository.NewGormRepository(db)
	return ui.NewModel(repo)
}

func main() {
	initModel := initialModel()
	p := tea.NewProgram(initModel)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
