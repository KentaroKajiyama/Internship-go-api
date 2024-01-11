package todo

import (
	"context"

	errDomain "github.com/KentaroKajiyama/Internship-go-api/domain/error"
	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

type DeleteToDoUseCase struct {
	todoRepository todoDomain.ToDoRepository
}

func NewDeleteToDoUseCase(todoRepository todoDomain.ToDoRepository) *DeleteToDoUseCase {
	return &DeleteToDoUseCase{todoRepository: todoRepository}
}

// todo項目削除
type DeleteToDoUseCaseInputDto struct {
	ID          string
	TodoID      string
	IsDeletable bool
}

// 新規項目を作成してリポジトリに登録する
func (uc *DeleteToDoUseCase) Delete(ctx context.Context, dto DeleteToDoUseCaseInputDto) error {
	todo, err := uc.todoRepository.Find(ctx, dto.ID, dto.TodoID)
	if err != nil {
		return err
	}
	if !dto.IsDeletable {
		return uc.todoRepository.Delete(ctx, todo)
	}
	return errDomain.NewError("削除保護が有効になっているため、削除できません。")
}
