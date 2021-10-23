package contract

type CartItem struct {
	Cart Cart `gorm:"foreignKey:CartId"`
	CartId uint `gorm:"primaryKey"`
	ItemId uint `gorm:"primaryKey"`
	Item Item `gorm:"foreignKey:ItemId"`
	Quantity int
}