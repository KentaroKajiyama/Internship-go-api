package todo

import (
	"context"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

type GetTagsInTodoUseCase struct {
	todoRepository todoDomain.TodoRepository
}

func NewGetTagsInTodoUseCase(todoRepository todoDomain.TodoRepository) *GetTagsInTodoUseCase {
	return &GetTagsInTodoUseCase{todoRepository: todoRepository}
}

type GetTagsInTodoUseCaseInputDto struct {
	TodoId string
}

func (uc *GetTagsInTodoUseCase) GetTagsInTodo(ctx context.Context, dto GetTagsInTodoUseCaseInputDto) (*todoDomain.TagsInTodo, error) {
	return uc.todoRepository.GetTagsInTodo(ctx, dto.TodoId)
}
