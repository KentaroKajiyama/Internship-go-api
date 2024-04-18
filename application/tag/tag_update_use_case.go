package tag

import (
	"context"

	tagDomain "github.com/KentaroKajiyama/Internship-go-api/domain/tag"
)

type UpdateTagUseCase struct {
	tagRepository tagDomain.TagRepository
}

func NewUpdateTagUseCase(tagRepository tagDomain.TagRepository) *UpdateTagUseCase {
	return &UpdateTagUseCase{tagRepository: tagRepository}
}

// tag項目更新
type UpdateTagUseCaseInputDto struct {
	Id    string
	TagId uint64
	Name  string
}

// 特定の項目を変更してリポジトリに登録する
func (uc *UpdateTagUseCase) Update(ctx context.Context, dto UpdateTagUseCaseInputDto) (*tagDomain.Tag, error) {
	user, err := tagDomain.NewTagFirst(dto.Id, dto.TagId, dto.Name)
	if err != nil {
		return nil, err
	}
	return uc.tagRepository.Update(ctx, user)
}
