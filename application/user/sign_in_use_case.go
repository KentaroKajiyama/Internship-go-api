package user

import (
	"context"
	"time"

	userDomain "github.com/KentaroKajiyama/internship-go-api/domain/user"
)

type RegistUserUseCase struct {
	userRepository userDomain.UserRepository
}

func NewRegistUserUseCase(userRepository userDomain.UserRepository) *RegistUserUseCase {
	return &RegistUserUseCase{userRepository: userRepository}
}

// ユーザー登録
type RegistUserUseCaseInputDto struct {
	name  string
	email string
}

func (uc *RegistUserUseCase) Register(ctx context.Context, dto RegistUserUseCaseInputDto) error {
	user, err := userDomain.NewUser(dto.name, dto.email, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return uc.userRepository.Create(ctx, user)
}
