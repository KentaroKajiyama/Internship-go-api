package model

import (
	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

func (s *Todo) ToDomainTodo() (*todoDomain.Todo, error) {
	//元々作成ずみ（データベースに存在していたデータ）の確認はタイムスタンプでいいのか？
	return todoDomain.NewTodo(s.Id, s.TodoId, s.Title, s.Description, s.IsDeletable, s.CreatedAt, s.UpdatedAt)
}

func NewTodoFromDomainTodo(todo *todoDomain.Todo) *Todo {
	return &Todo{
		Id:          todo.Id(),
		TodoId:      todo.TodoId(),
		Title:       todo.Title(),
		Description: todo.Description(),
		IsDeletable: todo.IsDeletable(),
	}
}
