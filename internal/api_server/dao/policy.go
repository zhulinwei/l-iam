package dao

import (
	"context"
	"l-iam/internal/api_server/model"

	"gorm.io/gorm"
)

type PolicyDao interface {
	Create(ctx context.Context, policy *model.Policy) error
}

type policyDao struct {
	db *gorm.DB
}

func NewPolicyDao(db *gorm.DB) PolicyDao {
	return &policyDao{db: db}
}

func (d *policyDao) Create(ctx context.Context, policy *model.Policy) error {
	return d.db.Save(policy).Error
}
