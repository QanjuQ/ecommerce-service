package contract

import "time"

type Cart struct {
	ID uint `gorm:"primaryKey"`
	IsPurchased bool `gorm:"default:false"`
	CreatedAt time.Time
}

func (c *Cart) Purchase(){
	c.IsPurchased = true
}
