package todo

import (
	"context"

	"github.com/google/uuid"
)

type TodoRepository interface {
	Find(ctx context.Context, id uuid.UUID, todo_id int) (*Todo, error)
	Create(ctx context.Context, Todo *Todo) error
	Update(ctx context.Context, Todo *Todo) error
	Delete(ctx context.Context, Todo *Todo) error
}
