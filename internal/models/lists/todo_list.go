package lists

type TodoList struct {
	ID    uint   `gorm:"primaryKey"`
	Index int    `gorm:"not null"`
	Title string `gorm:"not null"`
}

type Lists []*TodoList

func (a Lists) Len() int           { return len(a) }
func (a Lists) Swap(i, j int)      { (a)[i], (a)[j] = (a)[j], (a)[i] }
func (a Lists) Less(i, j int) bool { return (a)[i].Index < (a)[j].Index }
