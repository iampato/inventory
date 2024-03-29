// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/iampato/inventory/supplier/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	AddEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.SupplierService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{AddEndpoint: MakeAddEndpoint(s)}
	for _, m := range mdw["Add"] {
		eps.AddEndpoint = m(eps.AddEndpoint)
	}
	return eps
}
