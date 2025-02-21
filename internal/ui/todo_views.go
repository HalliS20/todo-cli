package ui

import (
	"fmt"
	ut "todo-cli/pkg/colorizer"
)

func (m *Model) renderListView() string {
	colorizer := ut.NewColorizer()
	s := colorizer.Colors[ut.Purple] + colorizer.FontStyles[ut.Bold] + *m.ParentName + colorizer.Commands[ut.Reset]
	space := 42
	for i := 0; i < space-len(*m.ParentName)-len(m.Views[m.ActiveOp].Header); i++ {
		s += " "
	}
	opColor := colorizer.Colors[m.Views[m.ActiveOp].OpColor]
	s += opColor + colorizer.FontStyles[ut.Bold] + m.Views[m.ActiveOp].Header + colorizer.Commands[ut.Reset]
	s += "\n\n"

	for i, choice := range m.ShownTodos {

		title := choice.Title
		Cursor := " " // no cursor
		if m.Cursor == i {
			Cursor = colorizer.Colors[ut.Pink] + colorizer.FontStyles[ut.Bold] + ">" + colorizer.Commands[ut.Reset] // cursor!
			title = colorizer.Colors[ut.ThickGreen] + colorizer.FontStyles[ut.Bold] + choice.Title + colorizer.Commands[ut.Reset]
		}

		checkPart := "   "
		if !choice.Dir {
			checked := " " // not selected
			if choice.Done {
				checked = "x" // selected!
			}
			checkPart = fmt.Sprintf("[%s]", checked)
		}

		s += fmt.Sprintf("%s %s %s\n", Cursor, checkPart, title)
	}

	return s
}
