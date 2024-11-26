// #cgo CFLAGS: -Wno-nullability-completeness
package main

import (
	"fmt"
	"os"
	"todo-cli/internal/models"
	"todo-cli/internal/repository"
	"todo-cli/internal/ui"
	"todo-cli/pkg"

	"github.com/charmbracelet/bubbletea"
)

func initialModel() tea.Model {
	db, err := pkg.OpenSqLiteDatabase("./db/todo.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	repo := repository.NewSQLiteRepository(db)
	todoList := repo.GetAll()

	return ui.Model{
		Todos:      todoList,
		Repo:       repo,
		ActiveView: models.List,
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
