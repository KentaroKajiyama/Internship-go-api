package tag

import (
	"time"
	"unicode/utf8"

	errDomain "github.com/KentaroKajiyama/Internship-go-api/domain/error"
	"github.com/KentaroKajiyama/Internship-go-api/pkg/uuid"
)

type Tag struct {
	id        string
	tag_id    string
	name      string
	createdAT time.Time
	updatedAt time.Time
}

func (s *Tag) Id() string {
	return s.id
}

func (s *Tag) TagId() string {
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

func newTag(id string, tag_id string, name string, createdAT time.Time, updatedAt time.Time) (*Tag, error) {
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

/* tag_idをどう決めていくか、とりあえず10にしている => uuidで生成する */
func NewTag(id string, name string, createdAt, updatedAt time.Time) (*Tag, error) {
	return newTag(
		id,
		uuid.NewUUID(),
		name,
		createdAt,
		updatedAt,
	)
}

func ReconstructTag(id string, tag_id string, name string, createdAt, updatedAt time.Time) (*Tag, error) {
	return newTag(
		id,
		tag_id,
		name,
		createdAt,
		updatedAt,
	)
}
