package todo

import (
	"context"
)

type TodoRepository interface {
	Find(ctx context.Context, id string, todo_id string) (*Todo, error)
	Create(ctx context.Context, Todo *Todo) error
	Update(ctx context.Context, Todo *Todo) error
	Delete(ctx context.Context, Todo *Todo) error
}
