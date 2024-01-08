package tag

import (
	"context"

	tagDomain "github.com/KentaroKajiyama/Internship-go-api/domain/tag"
)

type DeleteTagUseCase struct {
	tagRepository tagDomain.TagRepository
}

func NewDeleteTagUseCase(tagRepository tagDomain.TagRepository) *DeleteTagUseCase {
	return &DeleteTagUseCase{tagRepository: tagRepository}
}

// todo項目削除
type DeleteTagUseCaseInputDto struct {
	id     string
	tag_id int
}

// 特定の項目を削除する
func (uc *DeleteTagUseCase) Delete(ctx context.Context, dto DeleteTagUseCaseInputDto) error {
	tag, err := uc.tagRepository.Find(ctx, dto.id, dto.tag_id)
	if err != nil {
		return err
	}
	return uc.tagRepository.Delete(ctx, tag)
}
