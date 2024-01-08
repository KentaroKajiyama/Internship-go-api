package todo

import (
	"context"
	"time"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

type CreateTodoUseCase struct {
	todoRepository todoDomain.TodoRepository
}

func NewCreateTodoUseCase(todoRepository todoDomain.TodoRepository) *CreateTodoUseCase {
	return &CreateTodoUseCase{todoRepository: todoRepository}
}

// todo項目新規作成
type CreateTodoUseCaseInputDto struct {
	id           string
	title        string
	description  string
	is_deletable string
}

// 新規項目を作成してリポジトリに登録する、userはどうする？
func (uc *CreateTodoUseCase) Create(ctx context.Context, dto CreateTodoUseCaseInputDto) error {
	todo, err := todoDomain.NewTodo(dto.id, dto.title, dto.description, dto.is_deletable, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return uc.todoRepository.Create(ctx, todo)
}
