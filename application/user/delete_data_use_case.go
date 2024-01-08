package user

import (
	"context"
	"time"

	userDomain "github.com/KentaroKajiyama/Internship-go-api/domain/user"
)

type DeleteUserUseCase struct {
	userRepository userDomain.UserRepository
}

func NewDeleteUserUseCase(userRepository userDomain.UserRepository) *RegistUserUseCase {
	return &RegistUserUseCase{userRepository: userRepository}
}

// ユーザー登録
type DeleteUserUseCaseInputDto struct {
	id    string
	name  string
	email string
}

func (uc *UpdateUserUseCase) Delete(ctx context.Context, dto DeleteUserUseCaseInputDto) error {
	user, err := userDomain.ReconstructUser(dto.id, dto.name, dto.email, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return uc.userRepository.Delete(ctx, user)
}
