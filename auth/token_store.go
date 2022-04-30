package auth

type TokenStore interface {
	CreateToken(userId string) (*TokenDetails, error)
	GetToken(id string) (string, error)
	RemoveToken(id string) (int64, error)
}
