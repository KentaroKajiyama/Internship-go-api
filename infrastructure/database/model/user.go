package model

import (
	userDomain "github.com/KentaroKajiyama/internship-go-api/domain/user"
)

func (s *User) ToDomainUser() userDomain.User {
	return userDomain.NewUser(s.Name, s.Email, s.CreatedAt, s.UpdatedAt)
}

func (s *User) NewUserFromDomainUser(user *userDomain.User) User {
	return User{
		Id:        user.Id(),
		Name:      user.Name(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}
}
