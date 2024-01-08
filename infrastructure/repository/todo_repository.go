package repository

import (
	"context"

	"github.com/KentaroKajiyama/internship-go-api/domain/todo"
)

type todoRepository struct {
}

func NewTodoRepository() todo.TodoRepository {
	return &todoRepository{}
}

func (r *todoRepository) Create(ctx context.Context, todo *todo.Todo) error {
}

func (r *todoRepository) Update(ctx context.Context, todo *todo.Todo) error {
}

func (r *todoRepository) Delete(ctx context.Context, todo *todo.Todo) error {
}
