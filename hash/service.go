package hash

import (
	"fmt"
)

type Service interface {
	GenerateFromPassword(password string) (string, error)
	MatchPassword(hash string, password string) bool
}

type service struct {
	provider Provider
}

func NewService(provider Provider) (Service, error) {
	if provider == nil {
		return nil, fmt.Errorf("hash.NewService: provider is required")
	}

	return &service{
		provider,
	}, nil
}

func (svc *service) GenerateFromPassword(password string) (string, error) {
	hash, err := svc.provider.FromPassword(password)
	if err != nil {
		return "", fmt.Errorf("hash.Service: failed to generate hash from password (%s)", err)
	}

	return string(hash), nil
}

func (svc *service) MatchPassword(hash string, password string) bool {
	err := svc.provider.MatchPassword([]byte(hash), password)
	if err != nil {
		return false
	}

	return true
}
