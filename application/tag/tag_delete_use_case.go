package tag

import (
	"context"
	"time"

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
	ID        string
	TagID     int
	CreatedAt time.Time
}

// 特定の項目を削除する
func (uc *DeleteTagUseCase) Delete(ctx context.Context, dto DeleteTagUseCaseInputDto) error {
	tag, err := uc.tagRepository.Find(ctx, dto.ID, dto.TagID)
	if err != nil {
		return err
	}
	return uc.tagRepository.Delete(ctx, tag)
}
