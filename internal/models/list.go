package models

type TodoList struct {
	ID    uint   `gorm:"primaryKey"`
	Index int    `gorm:"not null"`
	Title string `gorm:"not null"`
}
