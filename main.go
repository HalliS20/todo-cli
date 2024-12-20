package main

import (
	"fmt"
	"os"
	"path/filepath"
	sqli "todo-cli/internal/repository/gormRepository"
	"todo-cli/internal/ui"
	"todo-cli/pkg"

	"github.com/charmbracelet/bubbletea"
)

func initialModel() tea.Model {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pathRest := ".todo/db/todo.db"
	dbPath := filepath.Join(homeDir, pathRest)

	db, err := pkg.OpenSqLiteDatabase(dbPath, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	repo := sqli.NewGormRepository(db)
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
