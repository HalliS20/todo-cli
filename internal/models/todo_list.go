package models

type TodoList struct {
	ID    uint   `gorm:"primaryKey"`
	Index int    `gorm:"not null"`
	Title string `gorm:"not null"`
}

type ListByIndex []TodoList

func (a ListByIndex) Len() int           { return len(a) }
func (a ListByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ListByIndex) Less(i, j int) bool { return a[i].Index < a[j].Index }
