package auth

import (
	"context"

	"github.com/KentaroKajiyama/Internship-go-api/domain"
)

type mock struct {
}

func (a mock) VerifyToken(ctx context.Context, token string) (string, error) {
	return token, nil
}

func NewAuthMockClient() domain.AuthClient {
	return &mock{}
}
