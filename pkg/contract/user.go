package contract

import "time"

type User struct {
	ID        uint
	Name      string
	Email     string `gorm:"unique"`
	Username  string `gorm:"primaryKey"`
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	CartId    uint
	Cart      Cart `gorm:"foreignKey:CartId"`
}
