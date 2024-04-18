package todo

import (
	"context"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

type AddTagsInTodoUseCase struct {
	todoRepository todoDomain.TodoRepository
}

func NewAddTagsInTodoUseCase(todoRepository todoDomain.TodoRepository) *AddTagsInTodoUseCase {
	return &AddTagsInTodoUseCase{todoRepository: todoRepository}
}

type AddTagsInTodoUseCaseInputDto struct {
	TodoId string
	TagIds []uint64
}

func (uc *AddTagsInTodoUseCase) AddTagsInTodo(ctx context.Context, dto AddTagsInTodoUseCaseInputDto) (*todoDomain.TagsInTodo, error) {
	// UseCase -> Infrastructure
	return uc.todoRepository.AddTagsInTodo(ctx, dto.TodoId, dto.TagIds)
}
