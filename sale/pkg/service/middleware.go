package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(SaleService) SaleService

type loggingMiddleware struct {
	logger log.Logger
	next   SaleService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a SaleService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next SaleService) SaleService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetReports(ctx context.Context) (results string, err error) {
	defer func() {
		l.logger.Log("method", "GetReports", "results", results, "err", err)
	}()
	return l.next.GetReports(ctx)
}
