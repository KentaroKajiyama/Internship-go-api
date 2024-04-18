package todo

import (
	"context"
)

type TodoRepository interface {
	Find(ctx context.Context, id string, todo_id string) (*Todo, error)
	FindMultiple(ctx context.Context, id string, todo_id string, title string) ([]*Todo, error)
	GetTagsInTodo(ctx context.Context, todo_id string) (*TagsInTodo, error)
	Create(ctx context.Context, Todo *Todo) (*Todo, error)
	Update(ctx context.Context, Todo *Todo) (*Todo, error)
	AddTagsInTodo(ctx context.Context, todo_id string, tag_id_s []uint64) (*TagsInTodo, error)
	Delete(ctx context.Context, Todo *Todo) (*Todo, error)
	DeleteMultiple(ctx context.Context, todos *DeleteTodosDto) (*DeleteTodosDto, error)
	DeleteTagsInTodo(ctx context.Context, todo_id string, tag_id_s []uint64) (*TagsInTodo, error)
}
