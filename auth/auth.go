package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

var ErrInvalidToken = errors.New("provided token is not valid")

type (
	// Claims specify JWT payload claims
	Claims struct {
		Subject   uint      `json:"sub"`
		ID        string    `json:"jti"`
		Issuer    string    `json:"iss"`
		Audience  string    `json:"aud"`
		IssuedAt  time.Time `json:"issued_at"`
		ExpiredAt time.Time `json:"exp"`
	}

	// TokenMaker is an interface for managing tokens
	TokenMaker interface {
		VerifyToken(token string) (*Claims, error)
		CreateToken(*Claims) (string, error)
	}

	// pasetoMaker is struct for paseto token maker
	pasetoMaker struct {
		paseto       *paseto.V2
		symmetricKey []byte
	}
)

func NewPasetoMaker(symmetricKey string) (TokenMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: symmetric key must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &pasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

// CreateToken implements TokenMaker.
// this method create a paseto token with specific payload
func (m *pasetoMaker) CreateToken(claims *Claims) (string, error) {
	return m.paseto.Encrypt(m.symmetricKey, claims, nil)
}

// VerifyToken implements TokenMaker.
// this method get token and decrypt token
func (m *pasetoMaker) VerifyToken(token string) (*Claims, error) {
	claims := &Claims{}

	err := m.paseto.Decrypt(token, m.symmetricKey, claims, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
