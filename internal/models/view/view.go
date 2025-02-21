package view

import (
	"todo-cli/pkg/colorizer"

	"github.com/charmbracelet/bubbletea"
)

type View struct {
	Update  func(msg tea.Msg) tea.Cmd
	View    func() string
	Footer  string
	Header  string
	OpColor colorizer.Colors
}
