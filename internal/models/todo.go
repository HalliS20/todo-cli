package models

type Todo struct {
	ID    uint   `gorm:"primaryKey"`
	Index int    `gorm:"not null"`
	Title string `gorm:"not null"`
	Done  bool   `gorm:"not null"`
}
