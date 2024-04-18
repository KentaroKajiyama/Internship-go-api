package user

import (
	"context"

	userDomain "github.com/KentaroKajiyama/Internship-go-api/domain/user"
)

type FindUserByUidUseCase struct {
	userRepository userDomain.UserRepository
}

func NewFindUserByUidUseCase(userRepository userDomain.UserRepository) *FindUserByUidUseCase {
	return &FindUserByUidUseCase{userRepository: userRepository}
}

// ユーザー登録
type FindUserByUidUseCaseInputDto struct {
	FirebaseUid string
}

func (uc *FindUserByUidUseCase) FindByUid(ctx context.Context, dto FindUserByUidUseCaseInputDto) (*userDomain.User, error) {
	// 名前とemailが被っていたら作れないようにしたい。
	return uc.userRepository.FindByUid(ctx, dto.FirebaseUid)
}
