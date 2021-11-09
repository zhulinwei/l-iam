package service

import (
	"context"
	"l-iam/internal/api_server/dao"
	"l-iam/internal/api_server/model"
)

type PolicyService interface {
	Create(ctx context.Context, policy *model.Policy) error
}

var _ PolicyService = (*policyService)(nil)

type policyService struct {
	dao dao.Factory
}

func newPolicies(dao dao.Factory) PolicyService {
	return &policyService{dao: dao}
}

func (s *policyService) Create(ctx context.Context, policy *model.Policy) error {
	return s.dao.Policies().Create(ctx, policy)
}
