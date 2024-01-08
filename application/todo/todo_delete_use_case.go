package todo

import (
	"context"

	"github.com/google/uuid"

	errDomain "github.com/KentaroKajiyama/internship-go-api/domain/error"
	todoDomain "github.com/KentaroKajiyama/internship-go-api/domain/todo"
)

type DeleteTodoUseCase struct {
	todoRepository todoDomain.TodoRepository
}

func NewDeleteTodoUseCase(todoRepository todoDomain.todoRepository) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{todoRepository: todoRepository}
}

// todo項目削除
type DeleteTodoUseCaseInputDto struct {
	id           uuid.UUID
	todo_id      int
	term_protect bool
}

// 新規項目を作成してリポジトリに登録する
func (uc *DeleteTodoUseCase) Deleteer(ctx context.Context, dto DeleteTodoUseCaseInputDto) error {
	todo, err := uc.todoRepository.Find(ctx, dto.id, dto.todo_id)
	if err != nil {
		return err
	}
	if !dto.term_protect {
		return uc.todoRepository.Delete(ctx, todo)
	}
	return errDomain.NewError("削除保護が有効になっているため、削除できません。")
}
