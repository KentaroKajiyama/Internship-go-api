package repository

import (
	"context"

	"github.com/KentaroKajiyama/internship-go-api/domain/user"
)

type userRepository struct {
}

func NewUserRepository() user.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(ctx context.Context, user *user.User) error {
}

func (r *userRepository) Update(ctx context.Context, user *user.User) error {
}

func (r *userRepository) Delete(ctx context.Context, user *user.User) error {
}
