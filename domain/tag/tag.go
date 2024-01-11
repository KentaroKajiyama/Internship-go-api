package tag

import (
	"time"
	"unicode/utf8"

	errDomain "github.com/KentaroKajiyama/Internship-go-api/domain/error"
	"github.com/KentaroKajiyama/Internship-go-api/pkg/uuid"
)

type Tag struct {
	id        string
	tagID     string
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func (s *Tag) ID() string {
	return s.id
}

func (s *Tag) TagID() string {
	return s.tagID
}

func (s *Tag) Name() string {
	return s.name
}

func (s *Tag) CreatedAt() time.Time {
	return s.createdAt
}

func (s *Tag) UpdatedAt() time.Time {
	return s.updatedAt
}

func newTag(id string, tagID string, name string, createdAt time.Time, updatedAt time.Time) (*Tag, error) {
	// バリデーション
	// idのバリデーション
	if !uuid.IsValid(id) {
		return nil, errDomain.NewError("UserIDが不正です。")
	}
	// ToDoIDのバリデーション
	if !uuid.IsValid(tagID) {
		return nil, errDomain.NewError("TagIDが不正です。")
	}
	// タイトルのバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("タイトルが不正です。")
	}
	return &Tag{
		id:        id,
		tagID:     tagID,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

const (
	// nameの最小値・最大値
	nameLengthMin = 1
	nameLengthMax = 255
)

/* tag_idをどう決めていくか、とりあえず10にしている => uuidで生成する */
func NewTag(id string, name string) (*Tag, error) {
	return newTag(
		id,
		uuid.NewUUID(),
		name,
		time.Now(),
		time.Now(),
	)
}

func ReconstructTag(id string, tagID string, name string, updatedAt time.Time) (*Tag, error) {
	return newTag(
		id,
		tagID,
		name,
		time.Now(),
		updatedAt,
	)
}
