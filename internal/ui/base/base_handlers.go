package ui_base

import (
	"unicode"

	"github.com/charmbracelet/bubbletea"
)

func HandleTextInput(msg tea.KeyMsg, textField *string) {
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

func HandleNavigation(msg tea.KeyMsg, curs *int, lisLen int) tea.Cmd {
	switch msg.String() {
	case "ctrl+c", "q": // quit
		return tea.Quit
	case "up", "k": // move cursor up
		if *curs > 0 {
			*curs--
		} else {
			*curs = lisLen - 1
		}
	case "down", "j": // move cursor down
		if *curs < lisLen-1 {
			*curs++
		} else {
			*curs = 0
		}
	}
	return nil
}
