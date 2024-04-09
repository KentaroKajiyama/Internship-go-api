package model

import (
	"time"

	userDomain "github.com/KentaroKajiyama/Internship-go-api/domain/user"
)

func (s *User) ToDomainUser() (*userDomain.User, error) {
	return userDomain.NewUser(s.Id, s.Name, s.Email, s.CreatedAt, s.UpdatedAt)
}

func NewUserFromDomainUser(user *userDomain.User) User {
	return User{
		Id:        user.Id(),
		Name:      user.Name(),
		Email:     user.Email(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
