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
	ID    string
	Name  string
	Email string
}

func (uc *UpdateUserUseCase) Delete(ctx context.Context, dto DeleteUserUseCaseInputDto) error {
	user, err := userDomain.ReconstructUser(dto.ID, dto.Name, dto.Email, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return uc.userRepository.Delete(ctx, user)
}
