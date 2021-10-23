package contract

import "time"

type Order struct {
	ID uint `gorm:"primaryKey"`
	CartId uint
	Cart Cart `gorm:"foreignKey:CartId"`
	UserId uint
	User User `gorm:"foreignKey:UserId"`
	CreatedAt time.Time
}