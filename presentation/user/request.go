package user

type GetUserParams struct {
	FirebaseUid string `param:"firebase_uid" validate:"required"`
}

type PostUsersParams struct {
	FirebaseUid string `json:"firebase_uid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
}

type PutUsersParams struct {
	Id    string `param:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type DeleteUsersParams struct {
	Id          string `param:"id" validate:"required"`
	FirebaseUid string `json:"firebase_uid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
}
