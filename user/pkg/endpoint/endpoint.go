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

// GetAllRolesRequest collects the request parameters for the GetAllRoles method.
type GetAllRolesRequest struct{}

// GetAllRolesResponse collects the response parameters for the GetAllRoles method.
type GetAllRolesResponse struct {
	Roles []entities.UserRole `json:"roles"`
}

// MakeGetAllRolesEndpoint returns an endpoint that invokes GetAllRoles on the service.
func MakeGetAllRolesEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		roles := s.GetAllRoles(ctx)
		return GetAllRolesResponse{Roles: roles}, nil
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

// GetAllUsersRequest collects the request parameters for the GetAllUsers method.
type GetAllUsersRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// GetAllUsersResponse collects the response parameters for the GetAllUsers method.
type GetAllUsersResponse struct {
	Users    []entities.UserModel `json:"users"`
	Total    int64                `json:"total"`
	Page     int64                `json:"page"`
	LastPage int64                `json:"last_page"`
}

// MakeGetAllUsersEndpoint returns an endpoint that invokes GetAllUsers on the service.
func MakeGetAllUsersEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllUsersRequest)
		users, total, page, lastPage := s.GetAllUsers(ctx, req.Limit, req.Offset)
		return GetAllUsersResponse{
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

// GetAllRoles implements Service. Primarily useful in a client.
func (e Endpoints) GetAllRoles(ctx context.Context) (roles []entities.UserRole) {
	request := GetAllRolesRequest{}
	response, err := e.GetAllRolesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllRolesResponse).Roles
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

// GetAllUsers implements Service. Primarily useful in a client.
func (e Endpoints) GetAllUsers(ctx context.Context, limit int, offset int) (users []entities.UserModel, total int64, page int64, lastPage int64) {
	request := GetAllUsersRequest{
		Limit:  limit,
		Offset: offset,
	}
	response, err := e.GetAllUsersEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllUsersResponse).Users, response.(GetAllUsersResponse).Total, response.(GetAllUsersResponse).Page, response.(GetAllUsersResponse).LastPage
}
