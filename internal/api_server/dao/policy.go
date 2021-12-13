package dao

import (
	"context"
	v1 "l-iam/internal/pkg/model/api_server/v1"
	"l-iam/pkg/code"
	"l-iam/pkg/errors"

	"gorm.io/gorm"
)

type PolicyDao interface {
	Create(ctx context.Context, policy *v1.Policy) error
	Get(ctx context.Context, policyId uint) (*v1.Policy, error)
	Update(ctx context.Context, policy *v1.Policy) error
	Delete(ctx context.Context, policyId uint) error
	List(ctx context.Context, opts *v1.PolicyQueryOptions) (*v1.PolicyList, error)
}

type policyDao struct {
	db *gorm.DB
}

func NewPolicyDao(db *gorm.DB) PolicyDao {
	return &policyDao{db: db}
}

func (d *policyDao) Create(ctx context.Context, policy *v1.Policy) error {
	return d.db.WithContext(ctx).Create(policy).Error
}

func (d *policyDao) Get(ctx context.Context, policyId uint) (*v1.Policy, error) {
	policy := &v1.Policy{}

	err := d.db.WithContext(ctx).Where("id = ?", policyId).First(policy).Error
	switch err {
	case nil:
		return policy, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, errors.WrapError(code.ErrUserNotFount, err)
	}
}

func (d *policyDao) Update(ctx context.Context, policy *v1.Policy) error {
	return d.db.WithContext(ctx).Updates(policy).Error
}

func (d *policyDao) Delete(ctx context.Context, policyId uint) error {
	return d.db.WithContext(ctx).Delete(&v1.Policy{}, policyId).Error
}

func (d *policyDao) List(ctx context.Context, opts *v1.PolicyQueryOptions) (*v1.PolicyList, error) {
	db := d.buildQueryDB(ctx, opts)
	res := &v1.PolicyList{}

	err := db.Find(&res.List).Offset(-1).Limit(-1).Count(&res.Total).Error
	switch err {
	case nil:
		return res, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, errors.WrapError(code.ErrUserNotFount, err)
	}
}

func (d *policyDao) buildQueryDB(ctx context.Context, opts *v1.PolicyQueryOptions) *gorm.DB {
	db := d.db.WithContext(ctx).Model(&v1.Policy{})
	if opts.Page > 0 && opts.Size > 0 {
		db = db.Offset((opts.Page - 1) * opts.Size).Limit(opts.Size)
	}
	if opts.Name != "" {
		db = db.Where("name like ?", "%"+opts.Name+"%")
	}
	return db.Order("id desc")
}
