package model

import (
	userDomain "github.com/KentaroKajiyama/Internship-go-api/domain/user"
)

func (s *User) ToDomainUser() (*userDomain.User, error) {
	if s.CreatedAt.IsZero() {
		return userDomain.NewUser(s.Name, s.Email)
	} else {
		return userDomain.ReconstructUser(s.ID, s.Name, s.Email, s.CreatedAt)
	}
}

func NewUserFromDomainUser(user *userDomain.User) User {
	return User{
		ID:        user.Id(),
		Name:      user.Name(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}
}
