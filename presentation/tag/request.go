package tag

type GetTagParams struct {
	Id    string `param:"id" validate:"required"`
	TagId string `param:"tag_id" validate:"required"`
	Name  string `query:"name"`
}

type GetTagsParams struct {
	Id     string `param:"id" json:"id" validate:"required"`
	TodoId string `query:"todo_id"`
	Name   string `query:"name"`
}

type PostTagsParams struct {
	Id   string `param:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type PutTagsParams struct {
	Id    string `param:"id" validate:"required"`
	TagId string `param:"tag_id" validate:"required"`
	Name  string `json:"name" validate:"required"`
}

type DeleteTagParams struct {
	Id    string `param:"id" validate:"required"`
	TagId string `param:"tag_id" validate:"required"`
}

type DeleteTagsParams struct {
	Id     string   `param:"id" validate:"required"`
	TagIds []string `json:"tag_id_s" validate:"required"`
}
