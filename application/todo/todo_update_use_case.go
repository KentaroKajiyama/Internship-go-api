package todo

import (
	"context"
	"time"

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
	id          string
	todo_id     int
	title       string
	description string
	isDeletable bool
}

// 特定の項目を変更してリポジトリに登録する
func (uc *UpdateTodoUseCase) Update(ctx context.Context, dto UpdateTodoUseCaseInputDto) error {
	todo, err := todoDomain.ReconstructTodo(dto.id, dto.todo_id, dto.title, dto.description, dto.isDeletable, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return uc.todoRepository.Update(ctx, todo)
}
