package models

type Todo struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"not null"`
	Done  bool   `gorm:"not null"`
}
