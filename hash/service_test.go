package hash

import (
	"fmt"
	"testing"
)

func TestServiceConstruction(t *testing.T) {
	t.Run("fails when provider is missing", func(t *testing.T) {
		if _, err := NewService(nil); err == nil {
			t.Fail()
		}
	})

	t.Run("returns a service with provider", func(t *testing.T) {
		provider := &mockProvider{}
		svc, _ := NewService(provider)
		if svc == nil {
			t.FailNow()
		}

		hashSvc, _ := svc.(*service)
		if hashSvc.provider != provider {
			t.Fail()
		}
	})
}

func TestServiceGenerationFromPassword(t *testing.T) {
	password := "a.very.strong.password"

	t.Run("fails when provider fails", func(t *testing.T) {
		svc, _ := NewService(&mockProvider{failOnGeneration: true})

		if _, err := svc.GenerateFromPassword(password); err == nil {
			t.Fail()
		}
	})

	t.Run("returns hashed password when all is well", func(t *testing.T) {
		svc, _ := NewService(&mockProvider{})

		hashedPassword, err := svc.GenerateFromPassword(password)
		if err != nil {
			t.Fail()
		}
		if hashedPassword == "" {
			t.Fail()
		}
	})
}

func TestServiceComparisonWithPassword(t *testing.T) {
	password := "a.very.strong.password"
	hash := "a.very.strong.password.hash"

	t.Run("fails when provider fails", func(t *testing.T) {
		svc, _ := NewService(&mockProvider{failOnComparison: true})

		if svc.MatchPassword(hash, password) {
			t.Fail()
		}
	})

	t.Run("succeeds when provider succeeds", func(t *testing.T) {
		svc, _ := NewService(&mockProvider{})

		if !svc.MatchPassword(hash, password) {
			t.Fail()
		}
	})
}

type mockProvider struct {
	failOnGeneration bool
	failOnComparison bool
}

func (mock *mockProvider) FromPassword(password string) ([]byte, error) {
	if mock.failOnGeneration {
		return nil, fmt.Errorf("failed to generate hash from password")
	}

	return []byte(password), nil
}

func (mock *mockProvider) MatchPassword(hash []byte, password string) error {
	if mock.failOnComparison {
		return fmt.Errorf("failed to match password with hash")
	}

	return nil
}
