package tag

import (
	"context"

	tagDomain "github.com/KentaroKajiyama/Internship-go-api/domain/tag"
)

type DeleteTagsUseCase struct {
	tagRepository tagDomain.TagRepository
}

func NewDeleteTagsUseCase(tagRepository tagDomain.TagRepository) *DeleteTagsUseCase {
	return &DeleteTagsUseCase{tagRepository: tagRepository}
}

type DeleteTagsUseCaseInputDto struct {
	Id     string
	TagIds []uint64
}

func (uc *DeleteTagsUseCase) DeleteTags(ctx context.Context, dto DeleteTagsUseCaseInputDto) (*tagDomain.Tags, error) {
	// UseCase -
	return uc.tagRepository.DeleteMultiple(ctx, dto.Id, dto.TagIds)
}
