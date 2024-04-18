package tag

import (
	"time"
	"unicode/utf8"

	errDomain "github.com/KentaroKajiyama/Internship-go-api/domain/error"
	"github.com/KentaroKajiyama/Internship-go-api/pkg/uuid"
)

type Tag struct {
	id        string
	tagId     uint64
	name      string
	createdAt time.Time
	updatedAt time.Time
}

type Tags struct {
	id     string
	tagIds []uint64
}

func (s *Tag) Id() string {
	return s.id
}

func (s *Tag) TagId() uint64 {
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

func (s *Tags) Id() string {
	return s.id
}

func (s *Tags) TagIds() []uint64 {
	return s.tagIds
}

func newTag(id string, tagId uint64, name string, createdAt time.Time, updatedAt time.Time) (*Tag, error) {
	// バリデーション
	// idのバリデーション
	if !uuid.IsValid(id) {
		return nil, errDomain.NewError("UserIDが不正です。")
	}
	// TagIDのバリデーション（必要になったら書く）

	// タイトルのバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("名前は32文字以内で設定してください。")
	}
	return &Tag{
		id:        id,
		tagId:     tagId,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func newTags(id string, tagIds []uint64) (*Tags, error) {
	// Validation
	if !uuid.IsValid(id) {
		return nil, errDomain.NewError("Idが不正です。")
	}
	if len(tagIds) == 0 {
		return nil, errDomain.NewError("TagIdが空です")
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
	return &Tags{
		id:     id,
		tagIds: tagIds,
	}, nil
}

const (
	// nameの最小値・最大値
	nameLengthMin = 1
	nameLengthMax = 32
)

/* tag_idをどう決めていくか→データベースで自動インクリメント*/
func NewTag(id string, tagId uint64, name string, createdAt time.Time, updatedAt time.Time) (*Tag, error) {
	return newTag(
		id,
		tagId,
		name,
		createdAt,
		updatedAt,
	)
}

func NewTagFirst(id string, tagId uint64, name string) (*Tag, error) {
	return newTag(
		id,
		tagId,
		name,
		time.Time{},
		time.Time{},
	)
}

func NewTagWithoutTime(id string, tagId uint64, name string) (*Tag, error) {
	return newTag(
		id,
		tagId,
		name,
		time.Time{},
		time.Time{},
	)
}

func NewTags(id string, tag_id_s []uint64) (*Tags, error) {
	return newTags(id, tag_id_s)
}
