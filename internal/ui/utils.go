package ui

import (
	"strconv"
	mo "todo-cli/internal/models/todo"
)

func ListWrap(lisa *[]*mo.Todo, up bool) {
	// if we go up we are moving from the end(0) to one from back (listEnd - 1)
	listEnd := len(*lisa) - 1

	end := 0
	oneFromOther := listEnd - 1

	// \/\/\/\/ Swap front and back \/\/\/\/ \\
	(*lisa)[0], (*lisa)[listEnd] = (*lisa)[listEnd], (*lisa)[0]
	(*lisa)[0].Index, (*lisa)[listEnd].Index = (*lisa)[listEnd].Index, (*lisa)[0].Index
	// now we have the one that we want to move at the other end
	// but the collateral item is at the other end when it should be next to the moved item

	if listEnd <= 1 {
		// if there are only 2 items we just need the swap
		return
	}
	if !up {
		// if we go down we are moving from the end(listEnd) to one from front (1)
		end = listEnd
		oneFromOther = 1
	}
	HoistItem(lisa, end, oneFromOther)
}

func HoistItem(lisa *[]*mo.Todo, from int, to int) {
	n := 1
	if to < from {
		n = -1
	}
	end := len(*lisa)
	for i := from; i != to; i += n {
		if i+n < 0 || i+n >= end || i < 0 || i >= end {
			panic("hoisting went out of bounds at i: " + strconv.Itoa(i) + " n:" + strconv.Itoa(n))
		}
		(*lisa)[i], (*lisa)[i+n] = (*lisa)[i+n], (*lisa)[i]
		(*lisa)[i].Index, (*lisa)[i+n].Index = (*lisa)[i+n].Index, (*lisa)[i].Index
	}
}
