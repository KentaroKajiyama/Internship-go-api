package todo

import (
	"context"

	errDomain "github.com/KentaroKajiyama/Internship-go-api/domain/error"
	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

type DeleteTodoUseCase struct {
	todoRepository todoDomain.TodoRepository
}

func NewDeleteTodoUseCase(todoRepository todoDomain.TodoRepository) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{todoRepository: todoRepository}
}

// todo項目削除
type DeleteTodoUseCaseInputDto struct {
	ID          string
	TodoID      int
	IsDeletable bool
}

// 新規項目を作成してリポジトリに登録する
func (uc *DeleteTodoUseCase) Delete(ctx context.Context, dto DeleteTodoUseCaseInputDto) error {
	todo, err := uc.todoRepository.Find(ctx, dto.ID, dto.TodoID)
	if err != nil {
		return err
	}
	if !dto.IsDeletable {
		return uc.todoRepository.Delete(ctx, todo)
	}
	return errDomain.NewError("削除保護が有効になっているため、削除できません。")
}
