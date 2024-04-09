package todo

import (
	"context"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

type FindTodoUseCase struct {
	todoRepository todoDomain.TodoRepository
}

func NewFindTodoUseCase(todoRepository todoDomain.TodoRepository) *FindTodoUseCase {
	return &FindTodoUseCase{todoRepository: todoRepository}
}

// todo項目検索
// データベース・フロントエンドからの受け皿
type FindTodoUseCaseInputDto struct {
	Id     string
	TodoId string
}

func (uc *FindTodoUseCase) Find(ctx context.Context, dto FindTodoUseCaseInputDto) (*todoDomain.Todo, error) {
	todo, err := uc.todoRepository.Find(ctx, dto.Id, dto.TodoId)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
