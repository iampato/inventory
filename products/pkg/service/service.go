package service

import "context"

// ProductsService describes the service.
type ProductsService interface {
	// Add your methods here
	Add(ctx context.Context, name string) error
}

type basicProductsService struct{}

func (b *basicProductsService) Add(ctx context.Context, name string) (e0 error) {
	// TODO implement the business logic of Add
	return e0
}

// NewBasicProductsService returns a naive, stateless implementation of ProductsService.
func NewBasicProductsService() ProductsService {
	return &basicProductsService{}
}

// New returns a ProductsService with all of the expected middleware wired in.
func New(middleware []Middleware) ProductsService {
	var svc ProductsService = NewBasicProductsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
