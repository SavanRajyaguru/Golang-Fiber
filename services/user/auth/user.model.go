package auth

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username  string  `json:"username" gorm:"not null;size:55" validate:"required"`
	Email     string  `json:"email" gorm:"not null" validate:"required,email"`
	Password  string  `json:"password" gorm:"not null" validate:"required,gte=6"`
	CompanyID uint    `json:"-"`
	Company   Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
}

func (UserModel) TableName() string {
	return "users"
}

type Company struct {
	gorm.Model
	UserInfo *[]UserModel `gorm:"foreignKey:CompanyID" json:"userInfo,omitempty"`
	Name     string
}

func (Company) TableName() string {
	return "company"
}
