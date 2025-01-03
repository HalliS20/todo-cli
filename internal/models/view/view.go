package view

import (
	"github.com/charmbracelet/bubbletea"
	"todo-cli/internal/enums/command"
	"todo-cli/pkg/colorizer"
)

type View struct {
	Update  func(msg tea.Msg) command.Command
	View    func() string
	Footer  string
	Header  string
	OpColor colorizer.Colors
}
