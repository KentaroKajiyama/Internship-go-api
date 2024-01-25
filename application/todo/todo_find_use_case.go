package todo

import (
	"context"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

type FindToDoUseCase struct {
	todoRepository todoDomain.ToDoRepository
}

func NewFindToDoUseCase(todoRepository todoDomain.ToDoRepository) *FindToDoUseCase {
	return &FindToDoUseCase{todoRepository: todoRepository}
}

// todo項目検索

type FindToDoUseCaseInputDto struct {
	ID     string
	ToDoID string
}

func (uc *FindToDoUseCase) Find(ctx context.Context, dto FindToDoUseCaseInputDto) (*todoDomain.ToDo, error) {
	todo, err := uc.todoRepository.Find(ctx, dto.ID, dto.ToDoID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
