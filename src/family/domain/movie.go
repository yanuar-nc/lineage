package domain

import "time"

type Family struct {
	ID          int       `gorm:"column:id; type:int; primaryKey;" json:"id"`
	Title       string    `gorm:"column:title; type:text;" json:"title"`
	Description string    `gorm:"column:description; type:text;" json:"description"`
	Status      bool      `gorm:"column:status; type:text;" json:"status"`
	Datetime    time.Time `gorm:"column:datetime; type:timestamp;" json:"datetime"`
}
