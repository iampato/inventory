package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "net/http"
	"strconv"

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

// makeFetchAllRolesHandler creates the handler logic
func makeFetchAllRolesHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/fetch-all-roles").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.FetchAllRolesEndpoint, decodeFetchAllRolesRequest, encodeFetchAllRolesResponse, options...)))
}

// decodeFetchAllRolesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeFetchAllRolesRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.FetchAllRolesRequest{}
	// err := json.NewDecoder(r.Body).Decode(&req)
	return req, nil
}

// encodeFetchAllRolesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeFetchAllRolesResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
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

// makeFetchAllUsersHandler creates the handler logic
func makeFetchAllUsersHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/fetch-all-users").Handler(
		handlers.CORS(handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.FetchAllUsersEndpoint, decodeFetchAllUsersRequest, encodeFetchAllUsersResponse, options...)))
}

// decodeFetchAllUsersRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeFetchAllUsersRequest(_ context.Context, r *http1.Request) (interface{}, error) {

	// get query params from r
	params := r.URL.Query()
	// params := mux.Vars(r) // how to get params from a UR
	var limitInt, pageInt int
	var err error
	limitStr := params.Get("limit")
	pageStr := params.Get("page")

	if len(limitStr) == 0 {
		limitInt = 15
	} else {
		limitInt, err = strconv.Atoi(limitStr)
		if err != nil {
			limitInt = 15
		}
	}

	if len(pageStr) == 0 {
		pageInt = 1
	} else {
		pageInt, err = strconv.Atoi(pageStr)
		if err != nil {
			pageInt = 1
		}
	}

	req := endpoint.FetchAllUsersRequest{
		Limit:  limitInt,
		Offset: pageInt,
	}
	// err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeFetchAllUsersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeFetchAllUsersResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
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
