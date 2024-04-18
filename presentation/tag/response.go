package tag

import "time"

type TagsResponseModel struct {
	Id        string    `json:"id"`
	TagId     string    `json:"tagId"`
	Name      string    `json:"name"`
	IsChecked bool      `json:"isChecked"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DeleteTagsResponseModel struct {
	Id     string   `json:"id"`
	TagIds []string `json:"tagIds"`
}
