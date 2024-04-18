package todo

import (
	"fmt"
	"reflect"
	"time"
	"unicode/utf8"

	errDomain "github.com/KentaroKajiyama/Internship-go-api/domain/error"
	"github.com/KentaroKajiyama/Internship-go-api/pkg/uuid"
)

type Todo struct {
	id          string
	todoId      string
	title       string
	description string
	isDeletable bool
	createdAt   time.Time
	updatedAt   time.Time
}
type TodosForDto struct {
	TodoId      string
	IsDeletable bool
}
type DeleteTodosDto struct {
	id    string
	todos []TodosForDto
}
type TagsInTodo struct {
	todoId string
	tagIds []uint64
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

func (s *DeleteTodosDto) Id() string {
	return s.id
}

func (s *DeleteTodosDto) Todos() []TodosForDto {
	return s.todos
}

func (s *TagsInTodo) TodoId() string {
	return s.todoId
}

func (s *TagsInTodo) TagIds() []uint64 {
	return s.tagIds
}

func newTodo(id string, todoId string, title string, description string, isDeletable bool, createdAt time.Time, updatedAt time.Time) (*Todo, error) {
	// バリデーション
	// idのバリデーション
	if !uuid.IsValid(id) {
		return nil, errDomain.NewError("UserIdが不正です。")
	}
	// ToDoIDのバリデーション
	if !uuid.IsValid(todoId) && todoId != "" {
		return nil, errDomain.NewError("TodoIdが不正です。")
	}
	// タイトルのバリデーション
	if utf8.RuneCountInString(title) < titleLengthMin || utf8.RuneCountInString(title) > titleLengthMax {
		return nil, errDomain.NewError("タイトルが不正です。255字以内にしてください。")
	}
	// 内容のバリデーション
	if utf8.RuneCountInString(description) < descriptionLengthMin || utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errDomain.NewError("内容が不正です。1000字以内にしてください。")
	}
	// 削除保護フラグのバリデーション
	if reflect.TypeOf(isDeletable).Kind() != reflect.Bool {
		return nil, errDomain.NewError("削除保護フラグが有効になっているため削除できません。")
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

func newDeleteTodosDto(id string, todos []TodosForDto) (*DeleteTodosDto, error) {
	if !uuid.IsValid(id) {
		return nil, errDomain.NewError("UserIdが不正です。")
	}
	if len(todos) == 0 {
		return nil, errDomain.NewError("Todosが空です。")
	}
	for _, todo := range todos {
		if !uuid.IsValid(todo.TodoId) {
			return nil, errDomain.NewError("TodoIdが不正です。")
		}
		if reflect.TypeOf(todo.IsDeletable).Kind() != reflect.Bool {
			return nil, errDomain.NewError("削除保護フラグが不正です。")
		}
	}
	return &DeleteTodosDto{
		id:    id,
		todos: todos,
	}, nil
}

func newTagsInTodo(todoId string, tagIds []uint64) (*TagsInTodo, error) {
	// Validation
	if !uuid.IsValid(todoId) {
		return nil, errDomain.NewError("TodoIdが不正です。")
	}
	if len(tagIds) == 0 {
		fmt.Printf("TagIdが空です\n")
		return &TagsInTodo{
			todoId: todoId,
			tagIds: tagIds,
		}, nil
	}
	seenTagIds := make(map[uint64]bool)
	for _, id := range tagIds {
		if id <= 0 {
			return nil, errDomain.NewError("TagIdが負の値です。")
		}
		if seenTagIds[id] {
			return nil, errDomain.NewError("TagIdが重複しています。")
		}
		seenTagIds[id] = true
	}
	return &TagsInTodo{
		todoId: todoId,
		tagIds: tagIds,
	}, nil
}

const (
	// Titleの最小値・最大値
	titleLengthMin = 0
	titleLengthMax = 255

	// 内容の最小値・最大値
	descriptionLengthMin = -1
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
func NewDeleteTodosDto(id string, todos []TodosForDto) (*DeleteTodosDto, error) {
	return newDeleteTodosDto(id, todos)
}
func NewTagsInTodo(todo_id string, tag_id_s []uint64) (*TagsInTodo, error) {
	return newTagsInTodo(todo_id, tag_id_s)
}
