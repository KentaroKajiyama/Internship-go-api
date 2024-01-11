package model

import (
	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

func (s *ToDo) ToDomainToDo() (*todoDomain.ToDo, error) {
	if s.CreatedAt.IsZero() {
		return todoDomain.NewToDo(s.ID, s.Title, s.Description, s.IsDeletable)
	} else {
		return todoDomain.ReconstructToDo(s.ID, s.ToDoID, s.Title, s.Description, s.IsDeletable, s.CreatedAt)
	}
}

func NewToDoFromDomainToDo(todo *todoDomain.ToDo) *ToDo {
	return &ToDo{
		ID:          todo.Id(),
		ToDoID:      todo.ToDoId(),
		Title:       todo.Title(),
		Description: todo.Description(),
		IsDeletable: todo.IsDeletable(),
		CreatedAt:   todo.CreatedAt(),
		UpdatedAt:   todo.UpdatedAt(),
	}
}
