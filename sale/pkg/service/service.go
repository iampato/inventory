package service

import "context"

// SaleService describes the service.
type SaleService interface {
	// Add your methods here
	GetReports(ctx context.Context) (results string, err error)
}

type basicSaleService struct{}

func (b *basicSaleService) GetReports(ctx context.Context) (results string, err error) {
	// TODO implement the business logic of GetReports
	return results, err
}

// NewBasicSaleService returns a naive, stateless implementation of SaleService.
func NewBasicSaleService() SaleService {
	return &basicSaleService{}
}

// New returns a SaleService with all of the expected middleware wired in.
func New(middleware []Middleware) SaleService {
	var svc SaleService = NewBasicSaleService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
