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

func NewTodoRepository(db *gorm.DB) todoDomain.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Find(ctx context.Context, id string, todo_id int) (*todoDomain.Todo, error) {
	conn := r.db.WithContext(ctx)
	var todoModel model.Todo
	var todoDomainPtr *todoDomain.Todo
	var errDom error
	if err := conn.Where("id = ? AND todo_id = ?", id, todo_id).Error; err != nil {
		return nil, err
	}
	todoDomainPtr, errDom = todoModel.ToDomainTodo()
	if errDom != nil {
		return nil, errDom
	}
	return todoDomainPtr, nil
}

func (r *todoRepository) Create(ctx context.Context, todo *todoDomain.Todo) error {
	conn := r.db.WithContext(ctx)
	todoModel := model.NewTodoFromDomainTodo(todo)
	if err := conn.Create(todoModel).Error; err != nil {
		return err
	}
	return nil
}

func (r *todoRepository) Update(ctx context.Context, todo *todoDomain.Todo) error {
	conn := r.db.WithContext(ctx)
	todoModel := model.NewTodoFromDomainTodo(todo)
	if err := conn.Update(todo.Id(), todoModel).Error; err != nil {
		return err
	}
	return nil
}

func (r *todoRepository) Delete(ctx context.Context, todo *todoDomain.Todo) error {
	conn := r.db.WithContext(ctx)
	todoModel := model.NewTodoFromDomainTodo(todo)
	if err := conn.Delete(todoModel).Error; err != nil {
		return err
	}
	return nil
}
