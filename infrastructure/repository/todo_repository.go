package repository

import (
	"context"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure/database/model"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(ctx context.Context, todo *todoDomain.Todo) error {
	conn := r.db.WithContext(ctx)
	todoModel := model.NewTodoFromDomainTodo(todo)
	if err := conn.Create(todoModel).Error(); err != nil {
		return err
	}
	return nil
}

func (r *todoRepository) Update(ctx context.Context, todo *todoDomain.Todo) error {
	conn := r.db.WithContext(ctx)
	todoModel := model.NewTodoFromDomainTodo(todo)
	if err := conn.Update(todoModel).Error(); err != nil {
		return err
	}
	return nil
}

func (r *todoRepository) Delete(ctx context.Context, todo *todoDomain.Todo) error {
	conn := r.db.WithContext(ctx)
	todoModel := model.NewTodoFromDomainTodo(todo)
	if err := conn.Delete(todoModel).Error(); err != nil {
		return err
	}
	return nil
}
