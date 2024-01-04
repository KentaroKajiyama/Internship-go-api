package tag

import (
	"time"
	//"kiravia.com/internship-go-api/domain/user"
)

type Tag struct {
	tag_id    int
	title     string
	createdAT time.Time
	updatedAt time.Time
}

func (s *Tag) Id() int {
	return s.tag_id
}

func (s *Tag) Title() string {
	return s.title
}

func (s *Tag) CreatedAt() time.Time {
	return s.createdAT
}

func (s *Tag) UpdatedAt() time.Time {
	return s.updatedAt
}

func newTag(tag_id int, title string, createdAT time.Time, updatedAt time.Time) (*Tag, error) {
	// バリデーション
	// タイトルのバリデーション
	// if utf8.RuneCountInString(title) < titleLengthMin && utf8.RuneCountInString(title) > titleLengthMax {
	// 	return nil, errDomain.NewError("タイトルが不正です。")
	// }
	return &Tag{
		tag_id:    tag_id,
		title:     title,
		createdAT: createdAT,
		updatedAt: updatedAt,
	}, nil
}

/*
const (
	// Titleの最小値・最大値
	titleLengthMin = 1
	titleLengthMax = 255
)
*/

/* tag_idをどう決めていくか、とりあえず10にしている */
func NewUser(title string, createdAt, updatedAt time.Time) (*Tag, error) {
	return newTag(
		10,
		title,
		createdAt,
		updatedAt,
	)
}

func ReconstructUser(tag_id int, title string, createdAt, updatedAt time.Time) (*Tag, error) {
	return newTag(
		tag_id,
		title,
		createdAt,
		updatedAt,
	)
}
