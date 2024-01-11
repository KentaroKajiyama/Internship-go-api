package todo

import (
	"reflect"
	"time"
	"unicode/utf8"

	errDomain "github.com/KentaroKajiyama/Internship-go-api/domain/error"
	"github.com/KentaroKajiyama/Internship-go-api/pkg/uuid"
)

// idは必要か？データベース的には必要だが、ビジネスロジック的には？
type ToDo struct {
	id          string
	todoID      string
	title       string
	description string
	isDeletable bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (s *ToDo) Id() string {
	return s.id
}

func (s *ToDo) ToDoId() string {
	return s.todoID
}

func (s *ToDo) Title() string {
	return s.title
}

func (s *ToDo) Description() string {
	return s.description
}

func (s *ToDo) IsDeletable() bool {
	return s.isDeletable
}

func (s *ToDo) CreatedAt() time.Time {
	return s.createdAt
}

func (s *ToDo) UpdatedAt() time.Time {
	return s.updatedAt
}

func newToDo(id string, todoID string, title string, description string, isDeletable bool, createdAt time.Time, updatedAt time.Time) (*ToDo, error) {
	// バリデーション
	// idのバリデーション
	if !uuid.IsValid(id) {
		return nil, errDomain.NewError("UserIDが不正です。")
	}
	// ToDoIDのバリデーション
	if !uuid.IsValid(todoID) {
		return nil, errDomain.NewError("ToDoIDが不正です。")
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
	return &ToDo{
		id:          id,
		todoID:      todoID,
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
func NewToDo(id string, title, description string, isDeletable bool) (*ToDo, error) {
	return newToDo(
		id,
		uuid.NewUUID(),
		title,
		description,
		isDeletable,
		time.Now(),
		time.Now(),
	)
}

func ReconstructToDo(id string, todoID string, title, description string, isDeletable bool, createdAt time.Time) (*ToDo, error) {
	return newToDo(
		id,
		todoID,
		title,
		description,
		isDeletable,
		createdAt,
		time.Now(),
	)
}
