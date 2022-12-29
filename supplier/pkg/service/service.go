package service

import "context"

// SupplierService describes the service.
type SupplierService interface {
	// Add your methods here
	Add(ctx context.Context, s string) (rs string, err error)
}

type basicSupplierService struct{}

func (b *basicSupplierService) Add(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Add
	return rs, err
}

// NewBasicSupplierService returns a naive, stateless implementation of SupplierService.
func NewBasicSupplierService() SupplierService {
	return &basicSupplierService{}
}

// New returns a SupplierService with all of the expected middleware wired in.
func New(middleware []Middleware) SupplierService {
	var svc SupplierService = NewBasicSupplierService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
