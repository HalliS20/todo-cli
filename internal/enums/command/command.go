package command

type Command int

const (
	Quit Command = iota
	None
	NavUp
	NavTo
)

func (c Command) String() string {
	switch c {
	case None:
		return "None"
	case NavTo:
		return "NavTo"
	case NavUp:
		return "NavUp"
	case Quit:
		return "Quit"
	}
	return "what this doesn't make ant sense"
}
