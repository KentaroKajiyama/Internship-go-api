package tag

import (
	"context"

	tagDomain "github.com/KentaroKajiyama/Internship-go-api/domain/tag"
)

type FindTagsUseCase struct {
	tagRepository tagDomain.TagRepository
}

func NewFindTagsUseCase(tagRepository tagDomain.TagRepository) *FindTagsUseCase {
	return &FindTagsUseCase{tagRepository: tagRepository}
}

type FindTagsUseCaseInputDto struct {
	Id    string
	TagId uint
	Name  string
}

func (uc *FindTagsUseCase) FindMultple(ctx context.Context, dto FindTagsUseCaseInputDto) ([]*tagDomain.Tag, error) {
	tags, err := uc.tagRepository.FindMultiple(ctx, dto.Id, dto.TagId, dto.Name)
	if err != nil {
		return nil, err
	}
	return tags, nil
}
