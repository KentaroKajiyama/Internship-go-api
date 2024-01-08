package model

import (
	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

func (s *Todo) ToDomainTodo() (*todoDomain.Todo, error) {
	return todoDomain.NewTodo(s.Id, s.Title, s.Text, s.IsDeletable, s.CreatedAt, s.UpdatedAt)
}

func NewTodoFromDomainTodo(todo *todoDomain.Todo) *Todo {
	return &Todo{
		Id:          todo.Id(),
		TodoId:      todo.TodoId(),
		Title:       todo.Title(),
		Text:        todo.Description(),
		IsDeletable: todo.IsDeletable(),
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}
}
