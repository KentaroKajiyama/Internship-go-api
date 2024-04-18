package user

import (
	"context"
)

type UserRepository interface {
	Find(ctx context.Context, id string) (*User, error)
	FindByUid(ctx context.Context, firebaseUid string) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, user *User) (*User, error)
}
