package tag

import (
	"context"
	"time"

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
	ID    string
	TagID int
	Title string
}

// 特定の項目を変更してリポジトリに登録する
func (uc *UpdateTagUseCase) Updateer(ctx context.Context, dto UpdateTagUseCaseInputDto) error {
	user, err := tagDomain.ReconstructTag(dto.ID, dto.TagID, dto.Title, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return uc.tagRepository.Update(ctx, user)
}
