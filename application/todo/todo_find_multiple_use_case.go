package todo

import (
	"context"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

type FindTodosUseCase struct {
	todoRepository todoDomain.TodoRepository
}

func NewFindTodosUseCase(repository todoDomain.TodoRepository) *FindTodosUseCase {
	return &FindTodosUseCase{todoRepository: repository}
}

type FindTodosUseCaseInputDto struct {
	Id     string
	TodoId string
	Title  string
}

func (uc *FindTodosUseCase) FindMultiple(ctx context.Context, dto FindTodosUseCaseInputDto) ([]*todoDomain.Todo, error) {
	todos, err := uc.todoRepository.FindMultiple(ctx, dto.Id, dto.TodoId, dto.Title)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
