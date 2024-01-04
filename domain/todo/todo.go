package todo

import (
	"time"
	//"kiravia.com/internship-go-api/domain/user"
	"unicode/utf8"

	errDomain "github.com/KentaroKajiyama/internship-go-api/domain/error"
)

type Todo struct {
	todo_id      int
	title        string
	description  string
	term_protect bool
	createdAT    time.Time
	updatedAt    time.Time
}

func (s *Todo) Id() int {
	return s.todo_id
}

func (s *Todo) Title() string {
	return s.title
}

func (s *Todo) Description() string {
	return s.description
}

func (s *Todo) TerminationProtect() bool {
	return s.term_protect
}

func (s *Todo) CreatedAt() time.Time {
	return s.createdAT
}

func (s *Todo) UpdatedAt() time.Time {
	return s.updatedAt
}

func newTodo(todo_id int, title string, description string, term_protect bool, createdAT time.Time, updatedAt time.Time) (*Todo, error) {
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
		todo_id:      todo_id,
		title:        title,
		description:  description,
		term_protect: term_protect,
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
func NewUser(title, description string, term_protect bool, createdAt, updatedAt time.Time) (*Todo, error) {
	return newTodo(
		10,
		title,
		description,
		term_protect,
		createdAt,
		updatedAt,
	)
}

func ReconstructUser(todo_id int, title, description string, term_protect bool, createdAt, updatedAt time.Time) (*Todo, error) {
	return newTodo(
		todo_id,
		title,
		description,
		term_protect,
		createdAt,
		updatedAt,
	)
}
