package todo

import "time"

type TodosResponseModel struct {
	Id          string    `json:"id"`
	TodoId      string    `json:"todoId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDeletable bool      `json:"isDeletable"`
	IsChecked   bool      `json:"isChecked"`
	TagIds      []string  `json:"tagIds"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
type PutTodosResponseModel struct {
	Id            string    `json:"id"`
	TodoId        string    `json:"todoId"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	IsDeletable   bool      `json:"isDeletable"`
	IsChecked     bool      `json:"isChecked"`
	AddedTagIds   []string  `json:"addedTagIds"`
	DeletedTagIds []string  `json:"deletedTagIds"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
type DeleteTodosResponseModel struct {
	Id    string `json:"id"`
	Todos []struct {
		TodoId      string `json:"todoId"`
		IsDeletable bool   `json:"isDeletable"`
	}
}
type TagsInTodoResponseModel struct {
	TodoId string   `json:"todoId"`
	TagIds []string `json:"tagIds"`
}
