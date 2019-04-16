package hash

import (
	"bytes"
	"testing"
)

func TestBCryptProviderCreate(t *testing.T) {
	t.Run("succeeds when all is well", func(t *testing.T) {
		if provider := NewBCryptProvider(); provider == nil {
			t.Fail()
		}
	})
}

func TestBCryptProviderHashGeneration(t *testing.T) {
	provider := NewBCryptProvider()
	password := "a.very.strong.password"

	t.Run("generates a hash from a password", func(t *testing.T) {
		hash, err := provider.FromPassword(password)
		if err != nil {
			t.Fail()
		}
		if hash == nil {
			t.Fail()
		}
	})

	t.Run("generates a unique hash every time", func(t *testing.T) {
		firstHash, _ := provider.FromPassword(password)
		secondHash, _ := provider.FromPassword(password)

		if bytes.Compare(firstHash, secondHash) == 0 {
			t.Fail()
		}
	})
}

func TestBCryptProviderHashComparison(t *testing.T) {
	provider := NewBCryptProvider()
	password := "a.very.strong.password"

	t.Run("compares a hash with its plain text version", func(t *testing.T) {
		hash, _ := provider.FromPassword(password)

		if err := provider.MatchPassword(hash, password); err != nil {
			t.Fail()
		}
	})
}
