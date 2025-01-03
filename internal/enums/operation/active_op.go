package operation

type ActiveOp int

const (
	Lister ActiveOp = iota
	Add
	Edit
	Empty
)
