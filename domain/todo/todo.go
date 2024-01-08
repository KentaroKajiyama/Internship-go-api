package todo

import (
	"time"
	//userDomain "github.com/KentaroKajiyama/internship-go-api/domain/user"
	"unicode/utf8"

	errDomain "github.com/KentaroKajiyama/internship-go-api/domain/error"
)

type Todo struct {
	id           string
	todo_id      int
	title        string
	description  string
	is_deletable bool
	createdAT    time.Time
	updatedAt    time.Time
}

func (s *Todo) Id() string {
	return s.id
}

func (s *Todo) TodoId() int {
	return s.todo_id
}

func (s *Todo) Title() string {
	return s.title
}

func (s *Todo) Description() string {
	return s.description
}

func (s *Todo) IsDeletable() bool {
	return s.is_deletable
}

func (s *Todo) CreatedAt() time.Time {
	return s.createdAT
}

func (s *Todo) UpdatedAt() time.Time {
	return s.updatedAt
}

func newTodo(id string, todo_id int, title string, description string, is_deletable bool, createdAT time.Time, updatedAt time.Time) (*Todo, error) {
	// バリデーション
	// タイトルのバリデーション
	if utf8.RuneCountInString(title) < titleLengthMin && utf8.RuneCountInString(title) > titleLengthMax {
		return nil, errDomain.NewError("タイトルが不正です。")
	}
	// 内容のバリデーション
	if utf8.RuneCountInString(description) < descriptionLengthMin && utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errDomain.NewError("内容が不正です。")
	}
	return &Todo{
		id:           id,
		todo_id:      todo_id,
		title:        title,
		description:  description,
		is_deletable: is_deletable,
		createdAT:    createdAT,
		updatedAt:    updatedAt,
	}, nil
}

const (
	// Titleの最小値・最大値
	titleLengthMin = 1
	titleLengthMax = 255

	// 内容の最小値・最大値
	descriptionLengthMin = 1
	descriptionLengthMax = 1000
)

/* Todo_idをどう決めていくか、とりあえず10にしている */
func NewTodo(id string, title, description string, is_deletable bool, createdAt, updatedAt time.Time) (*Todo, error) {
	return newTodo(
		id,
		10,
		title,
		description,
		is_deletable,
		createdAt,
		updatedAt,
	)
}

func ReconstructTodo(id string, todo_id int, title, description string, is_deletable bool, createdAt, updatedAt time.Time) (*Todo, error) {
	return newTodo(
		id,
		todo_id,
		title,
		description,
		is_deletable,
		createdAt,
		updatedAt,
	)
}
