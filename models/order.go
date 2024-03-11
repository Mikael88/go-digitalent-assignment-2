package models

import "time"

type Order struct {
	ID           uint       `gorm:"primaryKey" json:"orderId"`
	CustomerName string     `gorm:"type:varchar(255)" json:"customerName"`
	OrderedAt    time.Time  `json:"orderedAt"`
	Items        []Item     `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
	DeletedAt    *time.Time `gorm:"index" json:"-"`
}

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `gorm:"type:varchar(50)" json:"itemCode"`
	Description string `gorm:"type:text" json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `gorm:"index" json:"-"`
}
