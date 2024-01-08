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
	id     string
	tag_id int
	title  string
}

// 特定の項目を変更してリポジトリに登録する
func (uc *UpdateTagUseCase) Updateer(ctx context.Context, dto UpdateTagUseCaseInputDto) error {
	user, err := tagDomain.ReconstructTag(dto.id, dto.tag_id, dto.title, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return uc.tagRepository.Update(ctx, user)
}
