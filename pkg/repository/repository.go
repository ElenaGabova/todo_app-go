package repository;
type Autorization interface {
}

type UsersList interface {
}

type TodoItem interface {
}

type Repository struct {
	Autorization
	UsersList
	TodoItem
}
func NewRepository() *Repository{
	return &Repository{}
}