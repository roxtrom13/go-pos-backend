package dto

type Register struct {
	User     string `json:"user" form:"user"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
	Name     string `json:"name" form:"name" binding:"required"`
	LastName string `json:"last_name" form:"last_name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
}
