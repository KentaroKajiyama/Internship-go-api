package user

import (
	"context"

	userDomain "github.com/KentaroKajiyama/Internship-go-api/domain/user"
)

type UpdateUserUseCase struct {
	userRepository userDomain.UserRepository
}

func NewUpdateUserUseCase(userRepository userDomain.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepository: userRepository}
}

// ユーザー情報変更 この辺のuser情報の話はfirebaseを使っているのでフロントエンドだけかも
type UpdateUserUseCaseInputDto struct {
	Id    string
	Name  string
	Email string
}

func (uc *UpdateUserUseCase) Update(ctx context.Context, dto UpdateUserUseCaseInputDto) (*userDomain.User, error) {
	_, err := uc.userRepository.Find(ctx, dto.Id)
	if err != nil {
		return nil, err
	}
	//変更が必要かも
	user, err := userDomain.NewUserWithoutTime(dto.Id, dto.Name, dto.Email)
	if err != nil {
		return nil, err
	}
	return uc.userRepository.Update(ctx, user)
}
