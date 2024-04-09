package todo

// idは必要ない気がする。←認証まわり、jwtを実装できたらctxから取ってくる。todo_idもそんな感じがする。
type GetTodoParams struct {
	Id     string `param:"id" query:"id" json:"id" form:"id" validate:"required"`
	TodoId string `param:"todo_id" query:"todo_id" json:"todo_id" form:"todo_id"`
	Title  string `query:"title"`
}

type GetTodosParams struct {
	Id     string `param:"id" validate:"required"`
	TodoId string `json:"todo_id"`
	Title  string `query:"title" json:"title"`
}

type PostTodosParams struct {
	Id          string `param:"id" query:"id" json:"id" form:"id" validate:"required"`
	Title       string `json:"title" form:"title" query:"title" validate:"required"`
	Description string `json:"description" form:"description" query:"description"`
	IsDeletable bool   `json:"is_deletable" form:"is_deletable" query:"is_deletable" ` //ここでvalidate"required"にするとエラーが出る。
}

type PutTodosParams struct {
	Id          string `param:"id" query:"id" json:"id" form:"id" validate:"required"`
	TodoId      string `param:"todo_id" query:"todo_id" json:"todo_id" form:"todo_id" validate:"required"`
	Title       string `json:"title" form:"title" query:"title" validate:"required"`
	Description string `json:"description" form:"description" query:"description"`
	IsDeletable bool   `json:"is_deletable" form:"is_deletable" query:"is_deletable"`
}

type DeleteTodosParams struct {
	Id          string `param:"id" query:"id" json:"id" form:"id" validate:"required"`
	TodoId      string `param:"todo_id" query:"todo_id" json:"todo_id" form:"todo_id" validate:"required"`
	IsDeletable bool   `json:"is_deletable" form:"is_deletable" query:"is_deletable"`
}
