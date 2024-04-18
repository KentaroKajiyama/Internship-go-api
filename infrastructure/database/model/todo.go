package model

import (
	"fmt"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

func (s *Todo) ToDomainTodo() (*todoDomain.Todo, error) {
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

func ToDomainFromTodoTags(TodoTagModels []TodoTag) (*todoDomain.TagsInTodo, error) {
	var beforeTodoId string
	var todoId string
	var tagIds []uint64
	for index, tt := range TodoTagModels {
		todoId = tt.TodoId
		if todoId != beforeTodoId && index > 0 {
			return nil, fmt.Errorf("different todo ids are not allowed")
		}
		tagIds = append(tagIds, tt.TagId)
		beforeTodoId = todoId
	}
	return todoDomain.NewTagsInTodo(todoId, tagIds)
}

// Where should I check the uniqueness of tagIds? In my case, it has already been confirmed in domian layer.
func NewTodoTagModelArray(todoId string, tagIds []uint64) *[]TodoTag {
	tagMap := make(map[uint64]bool)
	var ttm TodoTag
	var ttms []TodoTag
	for _, tagId := range tagIds {
		if _, exists := tagMap[tagId]; !exists {
			ttm = TodoTag{
				TodoId: todoId,
				TagId:  tagId,
			}
			ttms = append(ttms, ttm)
		}
		tagMap[tagId] = true
	}
	return &ttms
}
