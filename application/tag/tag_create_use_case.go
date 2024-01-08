package tag

import (
	"context"
	"time"

	"github.com/google/uuid"

	tagDomain "github.com/KentaroKajiyama/internship-go-api/domain/tag"
)

type RegistTagUseCase struct {
	tagRepository tagDomain.tagRepository
}

func NewRegistTagUseCase(tagRepository tagDomain.tagRepository) *RegistTagUseCase {
	return &RegistTagUseCase{tagRepository: tagRepository}
}

// tag項目新規作成
type RegistTagUseCaseInputDto struct {
	id    uuid.UUID
	title string
}

// 新規項目を作成してリポジトリに登録する。
func (uc *RegistTagUseCase) Register(ctx context.Context, dto RegistTagUseCaseInputDto) error {
	user, err := tagDomain.NewTag(dto.title, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return uc.tagRepository.Create(ctx, user)
}
