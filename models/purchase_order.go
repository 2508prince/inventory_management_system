package models

import "time"

type PurchaseOrder struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int
	Supplier  string
	OrderDate time.Time
}
