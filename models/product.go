package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string `gorm:"type:varchar(255)" json:"name"`
	Description   string `gorm:"type:text" json:"description"`
	PurchasePrice uint   `json:"purchase_price"`
	SalePrice     uint   `json:"sale_price"`
	Stock         int16  `json:"stock"`
	Image         string `gorm:"type:varchar(255)" json:"image"`
	Unit          string `gorm:"type:varchar(50)" json:"unit"`
}
