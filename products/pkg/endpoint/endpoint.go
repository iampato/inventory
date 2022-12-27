package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/iampato/inventory/products/pkg/service"
)

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	Name string `json:"name"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	E0 error `json:"e0"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.ProductsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		e0 := s.Add(ctx, req.Name)
		return AddResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, name string) (e0 error) {
	request := AddRequest{Name: name}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).E0
}
