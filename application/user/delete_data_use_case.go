package user

import (
	"context"

	userDomain "github.com/KentaroKajiyama/Internship-go-api/domain/user"
)

type DeleteUserUseCase struct {
	userRepository userDomain.UserRepository
}

func NewDeleteUserUseCase(userRepository userDomain.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{userRepository: userRepository}
}

// ユーザー登録
type DeleteUserUseCaseInputDto struct {
	Id    string
	Name  string
	Email string
}

func (uc *DeleteUserUseCase) Delete(ctx context.Context, dto DeleteUserUseCaseInputDto) (*userDomain.User, error) {
	user, err := userDomain.NewUserWithoutTime(dto.Id, dto.Name, dto.Email)
	if err != nil {
		return nil, err
	}
	return uc.userRepository.Delete(ctx, user)
}
