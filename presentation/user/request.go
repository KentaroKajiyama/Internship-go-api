package user

type GetUsersParams struct {
	Id string `param:"id" query:"id" json:"id" form:"id" validate:"required"`
}

type PostUsersParams struct {
	Id    string `param:"id" query:"id" json:"id" form:"id" validate:"required"`
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required"`
}

type PutUsersParams struct {
	Id    string `param:"id" query:"id" json:"id" form:"id" validate:"required"`
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required"`
}

type DeleteUsersParams struct {
	Id    string `param:"id" query:"id" json:"id" form:"id" validate:"required"`
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required"`
}
