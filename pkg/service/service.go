package service;

type Autorization interface {
}

type UsersList interface {
}

type TodoItem interface {
}

type Service struct {
	Autorization
	UsersList
	TodoItem
}
func NewService(repos *repository.Repository) *Service{
	return &Service{}
}