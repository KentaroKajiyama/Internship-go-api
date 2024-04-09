package tag

type GetTagParams struct {
	Id    string `param:"id" query:"id" json:"id" form:"id"`
	TagId uint   `param:"tag_id" query:"tag_id" json:"tag_id" form:"tag_id"`
	Name  string `query:"name"`
}

type GetTagsParams struct {
	Id string `param:"id" json:"id" validate:"required"`
	TagId uint `json:"tag_id"`
	Name string `query:"name"`
}

type PostTagsParams struct {
	Id   string `param:"id" query:"id" json:"id" form:"id" `
	Name string `json:"name" form:"name" query:"name"`
}

type PutTagsParams struct {
	Id    string `param:"id" query:"id" json:"id" form:"id"`
	TagId uint   `param:"tag_id" query:"tag_id" json:"tag_id" form:"tag_id"`
	Name  string `json:"name" form:"name" query:"name"`
}

type DeleteTagsParams struct {
	Id    string `param:"id" query:"id" json:"id" form:"id"`
	TagId uint   `param:"tag_id" query:"tag_id" json:"tag_id" form:"tag_id"`
}
