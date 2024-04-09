package user

import (
	"time"
	"unicode/utf8"

	errDomain "github.com/KentaroKajiyama/Internship-go-api/domain/error"
	"github.com/KentaroKajiyama/Internship-go-api/pkg/uuid"
)

type User struct {
	id        string
	name      string
	email     string
	createdAt time.Time
	updatedAt time.Time
}

func (s *User) Id() string {
	return s.id
}

func (s *User) Name() string {
	return s.name
}

func (s *User) Email() string {
	return s.email
}

func (s *User) CreatedAt() time.Time {
	return s.createdAt
}

func (s *User) UpdatedAt() time.Time {
	return s.updatedAt
}

func newUser(id string, name string, email string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	// idのバリデーション
	if !uuid.IsValid(id) {
		return nil, errDomain.NewError("UserIDが不正です。")
	}
	// 名前のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin {
		return nil, errDomain.NewError("ユーザー名が不正です。")
	}
	// emailのバリデーション
	if utf8.RuneCountInString(email) < emailLengthMin {
		return nil, errDomain.NewError("emailが不正です。")
	}

	return &User{
		id:        id,
		name:      name,
		email:     email,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

const (
	// 名前の最小値
	nameLengthMin = 1

	// emailの最小値
	emailLengthMin = 1
)

func NewUser(id, name, email string, createdAt, updatedAt time.Time) (*User, error) {
	return newUser(
		id,
		name,
		email,
		createdAt,
		updatedAt,
	)
}

func NewUserWithoutTime(id string, name, email string) (*User, error) {
	return newUser(
		id,
		name,
		email,
		time.Time{},
		time.Time{},
	)
}

func NewUserWithoutIdAndTime(name, email string) (*User, error) {
	return newUser(
		uuid.NewUUID(),
		name,
		email,
		time.Time{},
		time.Time{},
	)
}
