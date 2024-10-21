package ui

import (
	"todo-cli/internal/models"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) renderEmptyView() string {
	col := models.NewColorizer()
	return col.Colors["pink"] + "Bye!" + col.Commands["reset"]
}

func (m Model) updateEmptyView(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}
