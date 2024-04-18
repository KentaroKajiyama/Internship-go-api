package user

import "time"

type UsersResponseModel struct {
	Id          string    `json:"id"`
	FirebaseUid string    `json:"firebaseUid"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
