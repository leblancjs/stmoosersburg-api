package hash

import "golang.org/x/crypto/bcrypt"

type bcryptProvider struct{}

func NewBCryptProvider() Provider {
	return &bcryptProvider{}
}

func (p *bcryptProvider) FromPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (p *bcryptProvider) MatchPassword(hash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}
