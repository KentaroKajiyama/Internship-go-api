package repository

import (
	"context"
	"fmt"

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
	// The value which is passed to Find should be pointer to struct or slice.
	if err := conn.Where("id = ? AND todo_id = ?", id, todo_id).Find(&todoModel).Error; err != nil {
		fmt.Printf("An error occurred")
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
	var todosModel *[]model.Todo
	var todosDomainPtr []*todoDomain.Todo
	// allocating memory for the slice
	todosModel = new([]model.Todo)

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
	if err := conn.Find(todosModel).Error; err != nil {
		return nil, err
	}
	//from infra to domain
	for _, tm := range *todosModel {
		td, err := tm.ToDomainTodo()
		if err != nil {
			return nil, err
		}
		todosDomainPtr = append(todosDomainPtr, td)
	}
	return todosDomainPtr, nil
}
func (r *todoRepository) GetTagsInTodo(ctx context.Context, todoId string) (*todoDomain.TagsInTodo, error) {
	conn := r.db.WithContext(ctx)
	var TodoTagsModels []model.TodoTag
	if err := conn.Where("todo_id=?", todoId).Find(&TodoTagsModels).Error; err != nil {
		return nil, err
	}
	if len(TodoTagsModels) == 0 {
		tagIds := make([]uint64, 0)
		return todoDomain.NewTagsInTodo(todoId, tagIds)
	}
	return model.ToDomainFromTodoTags(TodoTagsModels)
}
func (r *todoRepository) Create(ctx context.Context, todo *todoDomain.Todo) (*todoDomain.Todo, error) {
	conn := r.db.WithContext(ctx)
	//domain層からinfra層へ
	todoModel := model.NewTodoFromDomainTodo(todo)
	if err := conn.Create(todoModel).Error; err != nil {
		return nil, err
	}
	//データベース処理に問題がなければそのまま受け取ったtodo(domain)を返す。
	return todoModel.ToDomainTodo()
}

func (r *todoRepository) Update(ctx context.Context, todo *todoDomain.Todo) (*todoDomain.Todo, error) {
	conn := r.db.WithContext(ctx)
	//domain層からinfra層へ
	todoModel := model.NewTodoFromDomainTodo(todo)
	if err := conn.Model(model.Todo{}).Select("Title", "Description", "IsDeletable", "UpdatedAt").Where("id = ? AND todo_id = ?", todo.Id(), todo.TodoId()).Updates(&todoModel).Error; err != nil {
		return nil, err
	}
	//データベース処理に問題がなければそのまま受け取ったtodo(domain)を返す。
	return todoModel.ToDomainTodo()
}

func (r *todoRepository) AddTagsInTodo(ctx context.Context, todo_id string, tag_id_s []uint64) (*todoDomain.TagsInTodo, error) {
	if len(tag_id_s) == 0 {
		return todoDomain.NewTagsInTodo(todo_id, tag_id_s)
	}
	conn := r.db.WithContext(ctx)
	// From domain to infrastructure
	todoTagModel := model.NewTodoTagModelArray(todo_id, tag_id_s)
	// transaction?
	tx := conn.Begin()
	if err := conn.Create(todoTagModel).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return todoDomain.NewTagsInTodo(todo_id, tag_id_s)
}

func (r *todoRepository) Delete(ctx context.Context, todo *todoDomain.Todo) (*todoDomain.Todo, error) {
	conn := r.db.WithContext(ctx)
	//domain層からinfra層へ
	todoModel := model.NewTodoFromDomainTodo(todo)

	tx := conn.Begin()
	if err := conn.Where("id = ? AND todo_id = ?", todo.Id(), todo.TodoId()).Delete(todoModel).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if !todoModel.IsDeletable {
		tx.Rollback()
		return nil, fmt.Errorf("削除権限がありません。")
	}
	tx.Commit()
	//データベース処理に問題がなければそのまま受け取ったtodo(domain)を返す。
	return todoModel.ToDomainTodo()
}

func (r *todoRepository) DeleteMultiple(ctx context.Context, todos *todoDomain.DeleteTodosDto) (*todoDomain.DeleteTodosDto, error) {
	conn := r.db.WithContext(ctx)
	todo_ids := make([]string, 0)
	for _, todo := range todos.Todos() {
		todo_ids = append(todo_ids, todo.TodoId)
	}
	// From domain to infrastructure
	// transaction?
	// In gorm, you have to pass pointer to gorm methods.
	tx := conn.Begin()
	if err := conn.Where("id = ? AND todo_id IN(?)", todos.Id(), todo_ids).Delete(&model.Todo{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	// Return arguments directly if you don't have any errors.
	return todos, nil
}

func (r *todoRepository) DeleteTagsInTodo(ctx context.Context, todo_id string, tag_id_s []uint64) (*todoDomain.TagsInTodo, error) {
	if len(tag_id_s) == 0 {
		return todoDomain.NewTagsInTodo(todo_id, tag_id_s)
	}
	conn := r.db.WithContext(ctx)

	// transaction?
	// In gorm, you have to pass pointer to gorm methods.
	tx := conn.Begin()
	if err := conn.Where("todo_id = ? AND tag_id IN(?)", todo_id, tag_id_s).Delete(&model.TodoTag{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return todoDomain.NewTagsInTodo(todo_id, tag_id_s)
}
