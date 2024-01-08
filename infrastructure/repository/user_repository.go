package repository

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

/*
func (r *userRepository) Create(ctx context.Context, user *user.User) error {
}

func (r *userRepository) Update(ctx context.Context, user *user.User) error {
}

func (r *userRepository) Delete(ctx context.Context, user *user.User) error {
}
*/
