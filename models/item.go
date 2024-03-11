package models

type Item struct {
	ID          int64   `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `gorm:"type:varchar(50)" json:"itemCode"`
	Description string `gorm:"type:text" json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}
