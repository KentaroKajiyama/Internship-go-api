package tag

import (
	"time"
	"unicode/utf8"

	errDomain "github.com/KentaroKajiyama/Internship-go-api/domain/error"
	"github.com/KentaroKajiyama/Internship-go-api/pkg/uuid"
)

type Tag struct {
	id        string
	tagId     uint
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func (s *Tag) Id() string {
	return s.id
}

func (s *Tag) TagId() uint {
	return s.tagId
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

func newTag(id string, tagId uint, name string, createdAt time.Time, updatedAt time.Time) (*Tag, error) {
	// バリデーション
	// idのバリデーション
	if !uuid.IsValid(id) {
		return nil, errDomain.NewError("UserIDが不正です。")
	}
	// TagIDのバリデーション（必要になったら書く）

	// タイトルのバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("タイトルが不正です。")
	}
	return &Tag{
		id:        id,
		tagId:     tagId,
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

/* tag_idをどう決めていくか→データベースで自動インクリメント*/
func NewTag(id string, tagId uint, name string, createdAt time.Time, updatedAt time.Time) (*Tag, error) {
	return newTag(
		id,
		tagId,
		name,
		createdAt,
		updatedAt,
	)
}

func NewTagFirst(id string, tagId uint, name string) (*Tag, error) {
	return newTag(
		id,
		tagId,
		name,
		time.Time{},
		time.Time{},
	)
}

func NewTagWithoutTime(id string, tagId uint, name string) (*Tag, error) {
	return newTag(
		id,
		tagId,
		name,
		time.Time{},
		time.Time{},
	)
}

// func ReconstructTag(id string, tagID uint, name string) (*Tag, error) {
// 	return newTag(
// 		id,
// 		tagID,
// 		name,
// 	)
// }
