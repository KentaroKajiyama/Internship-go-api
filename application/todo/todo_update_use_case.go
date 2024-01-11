package todo

import (
	"context"
	"time"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

type UpdateToDoUseCase struct {
	todoRepository todoDomain.ToDoRepository
}

func NewUpdateToDoUseCase(todoRepository todoDomain.ToDoRepository) *UpdateToDoUseCase {
	return &UpdateToDoUseCase{todoRepository: todoRepository}
}

// todo項目更新
type UpdateToDoUseCaseInputDto struct {
	ID          string
	TodoID      string
	Title       string
	Description string
	IsDeletable bool
	CreatedAT   time.Time
}

// 特定の項目を変更してリポジトリに登録する
func (uc *UpdateToDoUseCase) Update(ctx context.Context, dto UpdateToDoUseCaseInputDto) error {
	todo, err := todoDomain.ReconstructToDo(dto.ID, dto.TodoID, dto.Title, dto.Description, dto.IsDeletable, dto.CreatedAT)
	if err != nil {
		return err
	}
	return uc.todoRepository.Update(ctx, todo)
}
