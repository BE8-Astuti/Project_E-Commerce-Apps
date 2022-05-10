package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID        uint
	Address       string `json:"address"`
	PaymentMethod string `json:"paymentMethod"`
	TotalBill     int    `json:"totalBill"`
	Status        string `json:"status"`
	Cart          []Cart `gorm:"foreignKey:OrderID;references:id"`
}
