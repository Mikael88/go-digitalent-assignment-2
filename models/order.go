package models

import "time"

type Order struct {
    ID           uint       `gorm:"primaryKey" json:"orderId"`
    CustomerName string     `gorm:"type:varchar(255)" json:"customerName"`
    OrderedAt    time.Time  `gorm:"column:ordered_at" json:"orderedAt"`
    Items        []Item     `gorm:"foreignKey:OrderID" json:"items"`
}