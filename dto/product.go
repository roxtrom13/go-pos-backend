package dto

type UpdateProduct struct {
	ID            uint   `json:"id" form:"id" binding:"required"`
	Name          string `json:"name" form:"name" binding:"required"`
	Description   string `json:"description" form:"description" binding:"required"`
	PurchasePrice uint   `json:"purchase_price" form:"purchase_price" binding:"required"`
	SalePrice     uint   `json:"sale_price" form:"sale_price" binding:"required"`
	Stock         int16  `json:"stock" form:"stock" binding:"required"`
	Image         string `json:"image" form:"image" binding:"required"`
	Unit          string `json:"unit" form:"unit" binding:"required"`
}

type CreateProduct struct {
	Name          string `json:"name" form:"name" binding:"required"`
	Description   string `json:"description" form:"description" binding:"required"`
	PurchasePrice uint   `json:"purchase_price" form:"purchase_price" binding:"required"`
	SalePrice     uint   `json:"sale_price" form:"sale_price" binding:"required"`
	Stock         int16  `json:"stock" form:"stock" binding:"required"`
	Image         string `json:"image" form:"image" binding:"required"`
	Unit          string `json:"unit" form:"unit" binding:"required"`
}
