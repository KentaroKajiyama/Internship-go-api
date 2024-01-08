package tag

import (
	"context"

	"github.com/google/uuid"
)

type DeleteTagUseCase struct {
	tagRepository tagDomain.tagRepository
}

func NewDeleteTagUseCase(tagRepository tagDomain.tagRepository) *DeleteTagUseCase {
	return &DeleteTagUseCase{tagRepository: tagRepository}
}

// todo項目削除
type DeleteTagUseCaseInputDto struct {
	id     uuid.UUID
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
