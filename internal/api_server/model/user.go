package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone    string `json:"phone" gorm:"column:phone" validate:"omitempty"`
	Email    string `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`
	Password string `json:"password,omitempty" gorm:"column:password" validate:"required"`
	Nickname string `json:"nickname" gorm:"column:nickname" validate:"required,min=1,max=30"`
}

type UserList struct {
	Total int    `json:"total"`
	List  []User `json:"list"`
}

func (*User) TableName() string {
	return "user"
}
