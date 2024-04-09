package todo

import (
	"reflect"
	"time"
	"unicode/utf8"

	errDomain "github.com/KentaroKajiyama/Internship-go-api/domain/error"
	"github.com/KentaroKajiyama/Internship-go-api/pkg/uuid"
)

// idは必要か？データベース的には必要だが、ビジネスロジック的には？
type Todo struct {
	id          string
	todoId      string
	title       string
	description string
	isDeletable bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (s *Todo) Id() string {
	return s.id
}

func (s *Todo) TodoId() string {
	return s.todoId
}

func (s *Todo) Title() string {
	return s.title
}

func (s *Todo) Description() string {
	return s.description
}

func (s *Todo) IsDeletable() bool {
	return s.isDeletable
}

func (s *Todo) CreatedAt() time.Time {
	return s.createdAt
}

func (s *Todo) UpdatedAt() time.Time {
	return s.updatedAt
}

func newTodo(id string, todoId string, title string, description string, isDeletable bool, createdAt time.Time, updatedAt time.Time) (*Todo, error) {
	// バリデーション
	// idのバリデーション
	if !uuid.IsValid(id) {
		return nil, errDomain.NewError("UserIdが不正です。")
	}
	// ToDoIDのバリデーション
	if !uuid.IsValid(todoId) {
		return nil, errDomain.NewError("TodoIdが不正です。")
	}
	// タイトルのバリデーション
	if utf8.RuneCountInString(title) < titleLengthMin || utf8.RuneCountInString(title) > titleLengthMax {
		return nil, errDomain.NewError("タイトルが不正です。")
	}
	// 内容のバリデーション
	if utf8.RuneCountInString(description) < descriptionLengthMin || utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errDomain.NewError("内容が不正です。")
	}
	// 削除保護フラグのバリデーション
	if reflect.TypeOf(isDeletable).Kind() != reflect.Bool {
		return nil, errDomain.NewError("削除保護フラグが不正です。")
	}
	return &Todo{
		id:          id,
		todoId:      todoId,
		title:       title,
		description: description,
		isDeletable: isDeletable,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
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

/* ToDo_idをどう決めていくか、とりあえず10にしている => uuidで生成することにする */
func NewTodo(id, todoId, title, description string, isDeletable bool, createdAt, updatedAt time.Time) (*Todo, error) {
	return newTodo(
		id,
		todoId,
		title,
		description,
		isDeletable,
		createdAt,
		updatedAt,
	)
}

func NewTodoWithoutTime(id, todoId, title, description string, isDeletable bool) (*Todo, error) {
	return newTodo(
		id,
		todoId,
		title,
		description,
		isDeletable,
		time.Time{},
		time.Time{},
	)
}

func NewTodoWithoutTodoIdAndTime(id, title, description string, isDeletable bool) (*Todo, error) {
	return newTodo(
		id,
		uuid.NewUUID(),
		title,
		description,
		isDeletable,
		time.Time{},
		time.Time{},
	)
}
