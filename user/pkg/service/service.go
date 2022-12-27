package service

import "context"

// UserService describes the service.
type UserService interface {
	// Add your methods here
	Create(ctx context.Context, name string) error
}

type basicUserService struct{}

func (b *basicUserService) Create(ctx context.Context, name string) (e0 error) {
	// TODO implement the business logic of Create
	return e0
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService() UserService {
	return &basicUserService{}
}

// New returns a UserService with all of the expected middleware wired in.
func New(middleware []Middleware) UserService {
	var svc UserService = NewBasicUserService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
