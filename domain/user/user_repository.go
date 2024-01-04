package user

type UserRepository interface {
	Create(User *User) (*User, error)
	Update(User *User) (*User, error)
}
