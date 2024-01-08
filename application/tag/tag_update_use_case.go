package tag

import (
	"context"
	"time"

	todoDomain "github.com/KentaroKajiyama/internship-go-api/domain/todo"
)

type ChangeTodoUseCase struct {
	todoRepository todoDomain.TodoRepository
}

func NewChangeTodoUseCase(todoRepository todoDomain.todoRepository) *ChangeTodoUseCase {
	return &ChangeTodoUseCase{todoRepository: todoRepository}
}

// todo項目更新
type ChangeTodoUseCaseInputDto struct {
	todo_id      int
	title        string
	description  string
	term_protect bool
}

// 特定の項目を変更してリポジトリに登録する
func (uc *ChangeTodoUseCase) Changeer(ctx context.Context, dto ChangeTodoUseCaseInputDto) error {
	user, err := todoDomain.ReconstructTodo(dto.todo_id, dto.title, dto.description, dto.term_protect, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return uc.todoRepository.Update(ctx, user)
}
