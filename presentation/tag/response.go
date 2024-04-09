package tag

import "time"

type TagsResponseModel struct {
	Id        string    `json:"id"`
	TagId     uint      `json:"tag_id"`
	Name      string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
