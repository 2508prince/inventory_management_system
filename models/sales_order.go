package models

import "time"

type SalesOrder struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int
	Total     float64
	Date      time.Time
}
