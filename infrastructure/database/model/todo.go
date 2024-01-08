package model

import (
	todoDomain "github.com/KentaroKajiyama/internship-go-api/domain/todo"
)

func (s *Todo) ToDomainTodo() todoDomain.Todo {
	return todoDomain.NewTodo(s.Id, s.Title, s.Text, s.IsDeletable, s.CreatedAt, s.UpdatedAt)
}

func (s *Todo) NewTodoFromDomainTodo(todo *todoDomain.Todo) *Todo {
	return &Todo{
		Id:          todo.Id(),
		TodoId:      todo.TodoId(),
		Title:       todo.Title(),
		Text:        todo.Text(),
		IsDeletable: todo.IsDeletable(),
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}
}
