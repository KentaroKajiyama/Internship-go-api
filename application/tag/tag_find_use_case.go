package tag

import (
	"context"

	tagDomain "github.com/KentaroKajiyama/Internship-go-api/domain/tag"
)

type FindTagUseCase struct {
	tagRepository tagDomain.TagRepository
}

func NewFindTagUseCase(tagRepository tagDomain.TagRepository) *FindTagUseCase {
	return &FindTagUseCase{tagRepository: tagRepository}
}

type FindTagUseCaseInputDto struct {
	Id    string
	TagId uint64
}

func (uc *FindTagUseCase) Find(ctx context.Context, dto FindTagUseCaseInputDto) (*tagDomain.Tag, error) {
	tag, err := uc.tagRepository.Find(ctx, dto.Id, dto.TagId)
	if err != nil {
		return nil, err
	}
	return tag, nil
}
