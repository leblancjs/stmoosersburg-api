package hash

type Provider interface {
	FromPassword(password string) ([]byte, error)
	MatchPassword(hash []byte, password string) error
}
