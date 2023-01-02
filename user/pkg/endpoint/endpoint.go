package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	entities "github.com/iampato/inventory/user/pkg/entities"
	service "github.com/iampato/inventory/user/pkg/service"
)

// CreateRoleRequest collects the request parameters for the CreateRole method.
type CreateRoleRequest struct {
	RoleName string `json:"role_name"`
}

// CreateRoleResponse collects the response parameters for the CreateRole method.
type CreateRoleResponse struct {
	Err error `json:"err"`
}

// MakeCreateRoleEndpoint returns an endpoint that invokes CreateRole on the service.
func MakeCreateRoleEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRoleRequest)
		err := s.CreateRole(ctx, req.RoleName)
		return CreateRoleResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r CreateRoleResponse) Failed() error {
	return r.Err
}

// FetchAllRolesRequest collects the request parameters for the FetchAllRoles method.
type FetchAllRolesRequest struct{}

// FetchAllRolesResponse collects the response parameters for the FetchAllRoles method.
type FetchAllRolesResponse struct {
	Roles []entities.UserRole `json:"roles"`
}

// MakeFetchAllRolesEndpoint returns an endpoint that invokes FetchAllRoles on the service.
func MakeFetchAllRolesEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		roles := s.FetchAllRoles(ctx)
		return FetchAllRolesResponse{Roles: roles}, nil
	}
}

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	RoleId       int    `json:"role_id"`
	Dob          string `json:"dob"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	User entities.UserModel `json:"user"`
	Err  error              `json:"err"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		user, err := s.Create(ctx, req.FirstName, req.LastName, req.EmailAddress, req.Phone, req.Password, req.RoleId, req.Dob)
		return CreateResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.Err
}

// FetchAllUsersRequest collects the request parameters for the FetchAllUsers method.
type FetchAllUsersRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// FetchAllUsersResponse collects the response parameters for the FetchAllUsers method.
type FetchAllUsersResponse struct {
	Users    []entities.UserModel `json:"users"`
	Total    int64                `json:"total"`
	Page     int64                `json:"page"`
	LastPage int64                `json:"last_page"`
}

// MakeFetchAllUsersEndpoint returns an endpoint that invokes FetchAllUsers on the service.
func MakeFetchAllUsersEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FetchAllUsersRequest)
		users, total, page, lastPage := s.FetchAllUsers(ctx, req.Limit, req.Offset)
		return FetchAllUsersResponse{
			LastPage: lastPage,
			Page:     page,
			Total:    total,
			Users:    users,
		}, nil
	}
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateRole implements Service. Primarily useful in a client.
func (e Endpoints) CreateRole(ctx context.Context, roleName string) (err error) {
	request := CreateRoleRequest{RoleName: roleName}
	response, err := e.CreateRoleEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateRoleResponse).Err
}

// FetchAllRoles implements Service. Primarily useful in a client.
func (e Endpoints) FetchAllRoles(ctx context.Context) (roles []entities.UserRole) {
	request := FetchAllRolesRequest{}
	response, err := e.FetchAllRolesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(FetchAllRolesResponse).Roles
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, firstName string, lastName string, emailAddress string, phone string, password string, roleId int, dob string) (user entities.UserModel, err error) {
	request := CreateRequest{
		Dob:          dob,
		EmailAddress: emailAddress,
		FirstName:    firstName,
		LastName:     lastName,
		Password:     password,
		Phone:        phone,
		RoleId:       roleId,
	}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).User, response.(CreateResponse).Err
}

// FetchAllUsers implements Service. Primarily useful in a client.
func (e Endpoints) FetchAllUsers(ctx context.Context, limit int, offset int) (users []entities.UserModel, total int64, page int64, lastPage int64) {
	request := FetchAllUsersRequest{
		Limit:  limit,
		Offset: offset,
	}
	response, err := e.FetchAllUsersEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(FetchAllUsersResponse).Users, response.(FetchAllUsersResponse).Total, response.(FetchAllUsersResponse).Page, response.(FetchAllUsersResponse).LastPage
}
