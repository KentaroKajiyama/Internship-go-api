package user

import (
	"context"
	"time"

	userDomain "github.com/KentaroKajiyama/internship-go-api/domain/user"
	"github.com/google/uuid"
)

type UpdateUserUseCase struct {
	userRepository userDomain.UserRepository
}

func NewUpdateUserUseCase(userRepository userDomain.UserRepository) *RegistUserUseCase {
	return &RegistUserUseCase{userRepository: userRepository}
}

// ユーザー情報変更
type UpdateUserUseCaseInputDto struct {
	id    uuid.UUID
	name  string
	email string
}

func (uc *UpdateUserUseCase) Update(ctx context.Context, dto UpdateUserUseCaseInputDto) error {
	user, err := userDomain.ReconstructUser(dto.id, dto.name, dto.email, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return uc.userRepository.Update(ctx, user)
}
