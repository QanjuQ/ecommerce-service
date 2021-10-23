package contract

import "time"

type Item struct {
	ID uint `gorm:"primaryKey"`
	Name string
	CreatedAt time.Time
}
