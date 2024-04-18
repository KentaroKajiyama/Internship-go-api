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
	Id    string
	TagId uint64
}

// 特定の項目を削除する
func (uc *DeleteTagUseCase) Delete(ctx context.Context, dto DeleteTagUseCaseInputDto) (*tagDomain.Tag, error) {
	tag, err := uc.tagRepository.Find(ctx, dto.Id, dto.TagId)
	if err != nil {
		return nil, err
	}
	return uc.tagRepository.Delete(ctx, tag)
}
