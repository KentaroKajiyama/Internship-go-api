package tag

import (
	"context"

	tagDomain "github.com/KentaroKajiyama/Internship-go-api/domain/tag"
)

type FindTagsByTodoIdUseCase struct {
	tagRepository tagDomain.TagRepository
}

func NewFindTagsByTodoIdUseCase(tagRepository tagDomain.TagRepository) *FindTagsByTodoIdUseCase {
	return &FindTagsByTodoIdUseCase{tagRepository: tagRepository}
}

type FindTagsByTodoIdUseCaseInputDto struct {
	Id     string
	TodoId string
	Name   string
}

func (uc *FindTagsByTodoIdUseCase) FindByTodoId(ctx context.Context, dto FindTagsByTodoIdUseCaseInputDto) ([]*tagDomain.Tag, error) {
	tags, err := uc.tagRepository.FindByTodoId(ctx, dto.Id, dto.TodoId, dto.Name)
	if err != nil {
		return nil, err
	}
	return tags, nil
}
