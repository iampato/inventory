package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/iampato/inventory/sale/pkg/service"
)

// GetReportsRequest collects the request parameters for the GetReports method.
type GetReportsRequest struct{}

// GetReportsResponse collects the response parameters for the GetReports method.
type GetReportsResponse struct {
	Results string `json:"results"`
	Err     error  `json:"err"`
}

// MakeGetReportsEndpoint returns an endpoint that invokes GetReports on the service.
func MakeGetReportsEndpoint(s service.SaleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		results, err := s.GetReports(ctx)
		return GetReportsResponse{
			Err:     err,
			Results: results,
		}, nil
	}
}

// Failed implements Failer.
func (r GetReportsResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetReports implements Service. Primarily useful in a client.
func (e Endpoints) GetReports(ctx context.Context) (results string, err error) {
	request := GetReportsRequest{}
	response, err := e.GetReportsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetReportsResponse).Results, response.(GetReportsResponse).Err
}
