package todo

type GetTodoParams struct {
	Id     string `param:"id" validate:"required"`
	TodoId string `query:"todo_id"`
}

type GetTodosParams struct {
	Id     string `param:"id" validate:"required"`
	TodoId string `query:"todo_id"`
	Title  string `query:"title"`
}
type GetTagsInTodoParams struct {
	TodoId string `param:"todo_id" validate:"required"`
}
type PostTagsInTodoParams struct {
	TodoId string   `parama:"todo_id" validate:"required"`
	TagIds []string `json:"tag_id_s" validate:"required"`
}

type DeleteTagsInTodoParams struct {
	TodoId string   `param:"todo_id" validate:"required"`
	TagIds []string `param:"tag_id_s" validate:"required"`
}
type PostTodosParams struct {
	Id          string   `param:"id" validate:"required"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description"`
	IsDeletable bool     `json:"is_deletable"`
	TagIds      []string `json:"tag_id_s" `
}

type PutTodosParams struct {
	Id           string   `param:"id" validate:"required"`
	TodoId       string   `param:"todo_id" validate:"required"`
	Title        string   `json:"title" validate:"required"`
	Description  string   `json:"description"`
	IsDeletable  bool     `json:"is_deletable"`
	PostTagIds   []string `json:"post_tag_id_s" `
	DeleteTagIds []string `json:"delete_tag_id_s"`
}

type DeleteTodoParams struct {
	Id          string `param:"id" validate:"required"`
	TodoId      string `param:"todo_id" validate:"required"`
	IsDeletable bool   `json:"is_deletable"`
}

type DeleteTodosParams struct {
	Id    string `param:"id" validate:"required"`
	Todos []struct {
		TodoId      string `json:"todo_id"`
		IsDeletable bool   `json:"is_deletable"`
	}
}
