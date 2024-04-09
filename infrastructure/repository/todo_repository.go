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

func (r *todoRepository) Find(ctx context.Context, id string, todo_id string) (*todoDomain.Todo, error) {
	conn := r.db.WithContext(ctx)
	var todoModel model.Todo
	var todoDomainPtr *todoDomain.Todo
	var errDom error
	if err := conn.Where("id = ? AND todo_id = ?", id, todo_id).Find(&todoModel).Error; err != nil {
		return nil, err
	}
	//infra層からdomain層へ
	todoDomainPtr, errDom = todoModel.ToDomainTodo()
	if errDom != nil {
		return nil, errDom
	}
	//上手くいったら取得したtodo(domain)を返す。
	return todoDomainPtr, nil
}

func (r *todoRepository) FindMultiple(ctx context.Context, id string, todo_id string, title string) ([]*todoDomain.Todo, error) {
	var todosModel []model.Todo
	var todosDomainPtr []*todoDomain.Todo
	conn := r.db.WithContext(ctx)
	if id != "" {
		conn = conn.Where("id = ?", id)
	}
	if todo_id != "" {
		conn = conn.Where("todo_id = ?", todo_id)
	}
	if title != "" {
		conn = conn.Where("title = ?", title)
	}
	if err := conn.Find(&todosModel).Error; err != nil {
		return nil, err
	}
	//from infra to domain
	for _, tm := range todosModel {
		td, err := tm.ToDomainTodo()
		if err != nil {
			return nil, err
		}
		todosDomainPtr = append(todosDomainPtr, td)
	}
	return todosDomainPtr, nil
}

func (r *todoRepository) Create(ctx context.Context, todo *todoDomain.Todo) (*todoDomain.Todo, error) {
	conn := r.db.WithContext(ctx)
	//domain層からinfra層へ
	todoModel := model.NewTodoFromDomainTodo(todo)
	if err := conn.Create(&todoModel).Error; err != nil {
		return nil, err
	}
	//データベース処理に問題がなければそのまま受け取ったtodo(domain)を返す。
	return todo, nil
}

func (r *todoRepository) Update(ctx context.Context, todo *todoDomain.Todo) (*todoDomain.Todo, error) {
	conn := r.db.WithContext(ctx)
	//domain層からinfra層へ
	todoModel := model.NewTodoFromDomainTodo(todo)
	if err := conn.Model(&model.Todo{}).Where("id = ? AND todo_id = ?", todo.Id(), todo.TodoId()).Updates(&todoModel).Error; err != nil {
		return nil, err
	}
	//データベース処理に問題がなければそのまま受け取ったtodo(domain)を返す。
	return todoModel.ToDomainTodo()
}

func (r *todoRepository) Delete(ctx context.Context, todo *todoDomain.Todo) (*todoDomain.Todo, error) {
	conn := r.db.WithContext(ctx)
	//domain層からinfra層へ
	todoModel := model.NewTodoFromDomainTodo(todo)
	if err := conn.Where("id = ? AND tag_id = ?", todo.Id(), todo.TodoId()).Delete(&todoModel).Error; err != nil {
		return nil, err
	}
	//データベース処理に問題がなければそのまま受け取ったtodo(domain)を返す。
	return todoModel.ToDomainTodo()
}
