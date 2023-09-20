package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	v5 "github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// GetAudience implements jwt.Claims.
func (pl *Payload) GetAudience() (v5.ClaimStrings, error) {
	return v5.ClaimStrings{pl.Username}, nil
}

// GetExpirationTime implements jwt.Claims.
func (pl *Payload) GetExpirationTime() (*v5.NumericDate, error) {
	return &v5.NumericDate{pl.ExpiredAt}, nil
}

// GetIssuedAt implements jwt.Claims.
func (pl *Payload) GetIssuedAt() (*v5.NumericDate, error) {
	return &v5.NumericDate{pl.IssuedAt}, nil
}

// GetIssuer implements jwt.Claims.
func (*Payload) GetIssuer() (string, error) {
	return "", nil
}

// GetNotBefore implements jwt.Claims.
func (*Payload) GetNotBefore() (*v5.NumericDate, error) {
	return nil, nil
}

// GetSubject implements jwt.Claims.
func (*Payload) GetSubject() (string, error) {
	return "", nil
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		ID:        uuid.New(),
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return fmt.Errorf("token is expired")
	}

	return nil
}
