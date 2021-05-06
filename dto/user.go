package dto

type UserUpdate struct {
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
	Name     string `json:"name" form:"name" binding:"required"`
	LastName string `json:"last_name" form:"last_name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
}

//type UserCreate struct {
	//Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
	//Name     string `json:"name" form:"name" binding:"required"`
	//LastName string `json:"last_name" form:"last_name" binding:"required"`
	//Email    string `json:"email" form:"email" binding:"required" validate:"email"`
//}
