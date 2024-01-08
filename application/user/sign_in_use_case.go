package user

import (
	"context"
	"time"

	userDomain "github.com/KentaroKajiyama/Internship-go-api/domain/user"
)

type RegistUserUseCase struct {
	userRepository userDomain.UserRepository
}

func NewRegistUserUseCase(userRepository userDomain.UserRepository) *RegistUserUseCase {
	return &RegistUserUseCase{userRepository: userRepository}
}

// ユーザー登録
type RegistUserUseCaseInputDto struct {
	Name  string
	Email string
}

func (uc *RegistUserUseCase) Register(ctx context.Context, dto RegistUserUseCaseInputDto) error {
	user, err := userDomain.NewUser(dto.Name, dto.Email, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return uc.userRepository.Create(ctx, user)
}
