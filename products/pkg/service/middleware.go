package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(ProductsService) ProductsService

type loggingMiddleware struct {
	logger log.Logger
	next   ProductsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a ProductsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next ProductsService) ProductsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Add(ctx context.Context, name string) (e0 error) {
	defer func() {
		l.logger.Log("method", "Add", "name", name, "e0", e0)
	}()
	return l.next.Add(ctx, name)
}
