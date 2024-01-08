package tag

import (
	"time"
	"unicode/utf8"

	errDomain "github.com/KentaroKajiyama/internship-go-api/domain/error"
	"github.com/google/uuid"
)

type Tag struct {
	id        uuid.UUID
	tag_id    int
	name      string
	createdAT time.Time
	updatedAt time.Time
}

func (s *Tag) Id() uuid.UUID {
	return s.id
}

func (s *Tag) TagId() int {
	return s.tag_id
}

func (s *Tag) Name() string {
	return s.name
}

func (s *Tag) CreatedAt() time.Time {
	return s.createdAT
}

func (s *Tag) UpdatedAt() time.Time {
	return s.updatedAt
}

func newTag(id uuid.UUID, tag_id int, name string, createdAT time.Time, updatedAt time.Time) (*Tag, error) {
	// バリデーション
	// タイトルのバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin && utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("タイトルが不正です。")
	}
	return &Tag{
		id:        id,
		tag_id:    tag_id,
		name:      name,
		createdAT: createdAT,
		updatedAt: updatedAt,
	}, nil
}

const (
	// nameの最小値・最大値
	nameLengthMin = 1
	nameLengthMax = 255
)

/* tag_idをどう決めていくか、とりあえず10にしている */
func NewTag(id uuid.UUID, name string, createdAt, updatedAt time.Time) (*Tag, error) {
	return newTag(
		id,
		10,
		name,
		createdAt,
		updatedAt,
	)
}

func ReconstructTag(id uuid.UUID, tag_id int, name string, createdAt, updatedAt time.Time) (*Tag, error) {
	return newTag(
		id,
		tag_id,
		name,
		createdAt,
		updatedAt,
	)
}
