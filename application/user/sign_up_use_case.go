package user

import (
	"context"

	userDomain "github.com/KentaroKajiyama/Internship-go-api/domain/user"
)

type SignUpUserUseCase struct {
	userRepository userDomain.UserRepository
}

func NewSignUpUserUseCase(userRepository userDomain.UserRepository) *SignUpUserUseCase {
	return &SignUpUserUseCase{userRepository: userRepository}
}

// ユーザー登録
type SignUpUserUseCaseInputDto struct {
	FirebaseUid string
	Name        string
	Email       string
}

func (uc *SignUpUserUseCase) SignUp(ctx context.Context, dto SignUpUserUseCaseInputDto) (*userDomain.User, error) {
	// 名前とemailが被っていたら作れないようにしたい。
	user, err := userDomain.NewUserWithoutIdAndTime(dto.FirebaseUid, dto.Name, dto.Email)
	if err != nil {
		return nil, err
	}
	return uc.userRepository.Create(ctx, user)
}
