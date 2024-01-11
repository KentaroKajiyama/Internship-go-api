package todo

import (
	"context"
)

type ToDoRepository interface {
	Find(ctx context.Context, id string, todo_id string) (*ToDo, error)
	Create(ctx context.Context, ToDo *ToDo) error
	Update(ctx context.Context, ToDo *ToDo) error
	Delete(ctx context.Context, ToDo *ToDo) error
}
