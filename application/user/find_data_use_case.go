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
	Id    string
	Name  string
	Email string
}

func (uc *FindUserUseCase) Find(ctx context.Context, dto FindUserUseCaseInputDto) (*userDomain.User, error) {
	// 名前とemailが被っていたら作れないようにしたい。
	user, err := userDomain.NewUserWithoutTime(dto.Id, dto.Name, dto.Email)
	if err != nil {
		return nil, err
	}
	return uc.userRepository.Find(ctx, user.Id())
}
