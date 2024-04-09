package todo

import (
	"context"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

type UpdateTodoUseCase struct {
	todoRepository todoDomain.TodoRepository
}

func NewUpdateTodoUseCase(todoRepository todoDomain.TodoRepository) *UpdateTodoUseCase {
	return &UpdateTodoUseCase{todoRepository: todoRepository}
}

// todo項目更新
type UpdateTodoUseCaseInputDto struct {
	Id          string
	TodoId      string
	Title       string
	Description string
	IsDeletable bool
}

// 特定の項目を変更してリポジトリに登録する
func (uc *UpdateTodoUseCase) Update(ctx context.Context, dto UpdateTodoUseCaseInputDto) (*todoDomain.Todo, error) {
	todo, err := todoDomain.NewTodoWithoutTime(dto.Id, dto.TodoId, dto.Title, dto.Description, dto.IsDeletable)
	if err != nil {
		return nil, err
	}
	return uc.todoRepository.Update(ctx, todo)
}
