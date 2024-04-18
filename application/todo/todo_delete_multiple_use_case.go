package todo

import (
	"context"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

type DeleteTodosUseCase struct {
	todoRepository todoDomain.TodoRepository
}

func NewDeleteTodosUseCase(todoRepository todoDomain.TodoRepository) *DeleteTodosUseCase {
	return &DeleteTodosUseCase{todoRepository: todoRepository}
}

type DeleteTodosUseCaseInputDto struct {
	Id    string
	Todos []todoDomain.TodosForDto
}

func (uc *DeleteTodosUseCase) DeleteTodos(ctx context.Context, dto DeleteTodosUseCaseInputDto) (*todoDomain.DeleteTodosDto, error) {
	todos, err := todoDomain.NewDeleteTodosDto(dto.Id, dto.Todos)
	if err != nil {
		return nil, err
	}
	return uc.todoRepository.DeleteMultiple(ctx, todos)
}
