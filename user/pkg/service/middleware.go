package service

import (
	"context"

	log "github.com/go-kit/kit/log"
	entities "github.com/iampato/inventory/user/pkg/entities"
)

// Middleware describes a service middleware.
type Middleware func(UserService) UserService

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserService) UserService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateRole(ctx context.Context, roleName string) (err error) {
	defer func() {
		l.logger.Log("method", "CreateRole", "roleName", roleName, "err", err)
	}()
	return l.next.CreateRole(ctx, roleName)
}
func (l loggingMiddleware) GetAllRoles(ctx context.Context) (roles []entities.UserRole) {
	defer func() {
		l.logger.Log("method", "GetAllRoles", "roles", roles)
	}()
	return l.next.GetAllRoles(ctx)
}
func (l loggingMiddleware) Create(ctx context.Context, firstName string, lastName string, emailAddress string, phone string, password string, roleId int, dob string) (user entities.UserModel, err error) {
	defer func() {
		l.logger.Log("method", "Create", "firstName", firstName, "lastName", lastName, "emailAddress", emailAddress, "phone", phone, "password", password, "roleId", roleId, "dob", dob, "user", user, "err", err)
	}()
	return l.next.Create(ctx, firstName, lastName, emailAddress, phone, password, roleId, dob)
}
func (l loggingMiddleware) GetAllUsers(ctx context.Context, limit int, offset int) (users []entities.UserModel, total int64, page int64, lastPage int64) {
	defer func() {
		l.logger.Log("method", "GetAllUsers", "limit", limit, "offset", offset, "users", users, "total", total, "page", page, "lastPage", lastPage)
	}()
	return l.next.GetAllUsers(ctx, limit, offset)
}
