package auth

import "context"

type AuthClient interface {
	VerifyToken(ctx context.Context, token string) (string, error)
}