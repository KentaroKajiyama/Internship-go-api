package todo

type TodoRepository interface {
	Create(Todo *Todo) (*Todo, error)
	Update(Todo *Todo) (*Todo, error)
}
