package models

type Items struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	ItemName string `gorm:"type:varchar(50)" json:"item_name"`
	Quantity int64  `json:"quantity"`
	Price    string `json:"price"`
}