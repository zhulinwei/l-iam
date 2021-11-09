package service

import "l-iam/internal/api_server/dao"

type Service interface {
	Policies() PolicyService
}

type service struct {
	dao dao.Factory
}

func NewService(dao dao.Factory) Service {
	return &service{dao: dao}
}

func (s *service) Policies() PolicyService {
	return newPolicies(s.dao)
}
