package dao

import "gorm.io/gorm"

type IPolicy interface {
}

type PolicyDao struct {
	db *gorm.DB
}

func NewPolicyDao(db *gorm.DB) IPolicy {
	return &PolicyDao{
		db: db,
	}
}
