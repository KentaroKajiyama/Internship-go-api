package user

import (
	"context"

	userDomain "github.com/KentaroKajiyama/Internship-go-api/domain/user"
)

type FindUserUseCase struct {
	userRepository userDomain.UserRepository
}

func NewFindUserUseCase(userRepository userDomain.UserRepository) *FindUserUseCase {
	return &FindUserUseCase{userRepository: userRepository}
}

// ユーザー登録
type FindUserUseCaseInputDto struct {
	Id string
}

func (uc *FindUserUseCase) Find(ctx context.Context, dto FindUserUseCaseInputDto) (*userDomain.User, error) {
	return uc.userRepository.Find(ctx, dto.Id)
}
