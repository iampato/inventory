package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/iampato/inventory/user/pkg/endpoint"
)

// makeCreateRoleHandler creates the handler logic
func makeCreateRoleHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/create-role").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CreateRoleEndpoint, decodeCreateRoleRequest, encodeCreateRoleResponse, options...)))
}

// decodeCreateRoleRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateRoleRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreateRoleRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateRoleResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateRoleResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetAllRolesHandler creates the handler logic
func makeGetAllRolesHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/get-all-roles").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetAllRolesEndpoint, decodeGetAllRolesRequest, encodeGetAllRolesResponse, options...)))
}

// decodeGetAllRolesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetAllRolesRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetAllRolesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetAllRolesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAllRolesResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateHandler creates the handler logic
func makeCreateHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/create").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CreateEndpoint, decodeCreateRequest, encodeCreateResponse, options...)))
}

// decodeCreateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetAllUsersHandler creates the handler logic
func makeGetAllUsersHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/get-all-users").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetAllUsersEndpoint, decodeGetAllUsersRequest, encodeGetAllUsersResponse, options...)))
}

// decodeGetAllUsersRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetAllUsersRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetAllUsersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetAllUsersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAllUsersResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
