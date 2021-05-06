package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Password string `gorm:"->;<-;not null" json:"-"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	LastName string `gorm:"type:varchar(255)" json:"last_name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
