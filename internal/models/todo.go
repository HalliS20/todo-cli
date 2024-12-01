package models

type Todo struct {
	ID    uint   `gorm:"primaryKey"`
	Index int    `gorm:"not null"`
	Title string `gorm:"not null"`
	Done  bool   `gorm:"not null"`
}

type ByIndex []Todo

func (a ByIndex) Len() int           { return len(a) }
func (a ByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIndex) Less(i, j int) bool { return a[i].Index < a[j].Index }
