package todo

import (
	"context"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

type DeleteTagsInTodoUseCase struct {
	todoRepository todoDomain.TodoRepository
}

func NewDeleteTagsInTodoUseCase(todoRepository todoDomain.TodoRepository) *DeleteTagsInTodoUseCase {
	return &DeleteTagsInTodoUseCase{todoRepository: todoRepository}
}

type DeleteTagsInTodoUseCaseInputDto struct {
	TodoId string
	TagIds []uint64
}

func (uc *DeleteTagsInTodoUseCase) DeleteTagsInTodo(ctx context.Context, dto DeleteTagsInTodoUseCaseInputDto) (*todoDomain.TagsInTodo, error) {
	// UseCase -
	return uc.todoRepository.DeleteTagsInTodo(ctx, dto.TodoId, dto.TagIds)
}
