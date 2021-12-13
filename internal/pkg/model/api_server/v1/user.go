package v1

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Username  string         `json:"username,omitempty" gorm:"column:username" validate:"required,min=1,max=30"`
	Password  string         `json:"password,omitempty" gorm:"column:password" validate:"required"`
	Phone     string         `json:"phone,omitempty" gorm:"column:phone" validate:"omitempty"`
	Email     string         `json:"email,omitempty" gorm:"column:email" validate:"required,email,min=1,max=100"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"deleted_at"`
}

func (*User) TableName() string {
	return "user"
}

type UserQueryUri struct {
	UserId uint `uri:"user_id"`
}

type UserQueryOptions struct {
	Name string `form:"name"`
	Page int    `form:"page" binding:"required,number,gt=0"`
	Size int    `form:"size" binding:"required,number,gt=0"`
}

type UserList struct {
	Total int64   `json:"total"`
	List  []*User `json:"list"`
}

func (u *User) Validate() error {
	return nil
}
