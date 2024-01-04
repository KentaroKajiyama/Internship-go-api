package user

import (
	"context"
	"time"

	userDomain "github.com/KentaroKajiyama/internship-go-api/domain/user"
	"github.com/google/uuid"
)

type ChangeUserUseCase struct {
	userRepository userDomain.UserRepository
}

func NewChangeUserUseCase(userRepository userDomain.UserRepository) *RegistUserUseCase {
	return &RegistUserUseCase{userRepository: userRepository}
}

// ユーザー登録
type ChangeUserUseCaseInputDto struct {
	id    uuid.UUID
	name  string
	email string
}

func (uc *ChangeUserUseCase) Changer(ctx context.Context, dto ChangeUserUseCaseInputDto) error {
	user, err := userDomain.ReconstructUser(dto.id, dto.name, dto.email, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return uc.userRepository.Update(ctx, user)
}
