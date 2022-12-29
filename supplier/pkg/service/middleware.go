package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(SupplierService) SupplierService

type loggingMiddleware struct {
	logger log.Logger
	next   SupplierService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a SupplierService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next SupplierService) SupplierService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Add(ctx context.Context, s string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Add", "s", s, "rs", rs, "err", err)
	}()
	return l.next.Add(ctx, s)
}
