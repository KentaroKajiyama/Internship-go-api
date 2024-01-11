package tag

import (
	"context"

	tagDomain "github.com/KentaroKajiyama/Internship-go-api/domain/tag"
)

type RegistTagUseCase struct {
	tagRepository tagDomain.TagRepository
}

func NewRegistTagUseCase(tagRepository tagDomain.TagRepository) *RegistTagUseCase {
	return &RegistTagUseCase{tagRepository: tagRepository}
}

// tag項目新規作成
type RegistTagUseCaseInputDto struct {
	ID    string
	Title string
}

// 新規項目を作成してリポジトリに登録する。
func (uc *RegistTagUseCase) Register(ctx context.Context, dto RegistTagUseCaseInputDto) error {
	user, err := tagDomain.NewTag(dto.ID, dto.Title)
	if err != nil {
		return err
	}
	return uc.tagRepository.Create(ctx, user)
}
