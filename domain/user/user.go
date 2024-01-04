package user

import (
	"time"
	// "unicode/utf8"

	"github.com/google/uuid"
	// errDomain "kiravia.com/internship-go-api/domain/error"
)

type User struct {
	id        uuid.UUID
	name      string
	email     string
	createdAt time.Time
	updatedAt time.Time
}

func (s *User) Id() uuid.UUID {
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

func newUser(id uuid.UUID, name string, email string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	// IDのバリデーション（必要ないかも）
	// if err := uuid.Validate(id); err != nil {
	// 	return nil, err
	// }
	// 名前のバリデーション
	// if utf8.RuneCountInString(name) < nameLengthMin {
	// 	return nil, errDomain.NewError("ユーザー名が不正です。")
	// }
	// emailのバリデーション

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

func NewUser(name, email string, createdAt, updatedAt time.Time) (*User, error) {
	return newUser(
		uuid.New(),
		name,
		email,
		createdAt,
		updatedAt,
	)
}

func ReconstructUser(id uuid.UUID, name, email string, createdAt, updatedAt time.Time) (*User, error) {
	return newUser(
		id,
		name,
		email,
		createdAt,
		updatedAt,
	)
}