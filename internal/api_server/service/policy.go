package service

import (
	"context"
	"l-iam/internal/api_server/dao"
	v1 "l-iam/internal/pkg/model/api_server/v1"
)

type PolicyService interface {
	Create(ctx context.Context, policy *v1.Policy) error
	Get(ctx context.Context, policyId uint) (*v1.Policy, error)
	Update(ctx context.Context, policyId uint, update *v1.Policy) error
	Delete(ctx context.Context, policyId uint) error
	List(ctx context.Context, opts *v1.PolicyQueryOptions) (*v1.PolicyList, error)
}

var _ PolicyService = (*policyService)(nil)

type policyService struct {
	dao dao.Factory
}

func newPolicyService(dao dao.Factory) PolicyService {
	return &policyService{dao: dao}
}

func (s *policyService) Create(ctx context.Context, policy *v1.Policy) error {
	return s.dao.Policies().Create(ctx, policy)
}

func (s *policyService) Get(ctx context.Context, policyId uint) (*v1.Policy, error) {
	return s.dao.Policies().Get(ctx, policyId)
}

func (s *policyService) Update(ctx context.Context, policyId uint, update *v1.Policy) error {
	policy, err := s.Get(ctx, policyId)
	if err != nil {
		return err
	}
	if update.Name != "" {
		policy.Name = update.Name
	}
	if policy.Policy != nil {
		policy.Policy = update.Policy
	}

	return s.dao.Policies().Update(ctx, policy)
}

func (s *policyService) List(ctx context.Context, opts *v1.PolicyQueryOptions) (*v1.PolicyList, error) {
	return s.dao.Policies().List(ctx, opts)
}

func (s *policyService) Delete(ctx context.Context, policyId uint) error {
	return s.dao.Policies().Delete(ctx, policyId)
}
