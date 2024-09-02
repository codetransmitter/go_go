// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// User defines model for User.
type User struct {
	CreatedAt *string `json:"createdAt,omitempty"`
	DeletedAt *string `json:"deletedAt,omitempty"`
	Email     *string `json:"email,omitempty"`
	Id        *uint   `json:"id,omitempty"`
	Password  *string `json:"password,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// DeleteUserJSONRequestBody defines body for DeleteUser for application/json ContentType.
type DeleteUserJSONRequestBody = User

// PatchUserJSONRequestBody defines body for PatchUser for application/json ContentType.
type PatchUserJSONRequestBody = User

// PostUserJSONRequestBody defines body for PostUser for application/json ContentType.
type PostUserJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Delete user by id
	// (DELETE /user)
	DeleteUser(ctx echo.Context) error
	// Get all users
	// (GET /user)
	GetUser(ctx echo.Context) error
	// Update user by id
	// (PATCH /user)
	PatchUser(ctx echo.Context) error
	// Create a new user
	// (POST /user)
	PostUser(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// DeleteUser converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUser(ctx)
	return err
}

// GetUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUser(ctx)
	return err
}

// PatchUser converts echo context to params.
func (w *ServerInterfaceWrapper) PatchUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchUser(ctx)
	return err
}

// PostUser converts echo context to params.
func (w *ServerInterfaceWrapper) PostUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUser(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE(baseURL+"/user", wrapper.DeleteUser)
	router.GET(baseURL+"/user", wrapper.GetUser)
	router.PATCH(baseURL+"/user", wrapper.PatchUser)
	router.POST(baseURL+"/user", wrapper.PostUser)

}

type DeleteUserRequestObject struct {
	Body *DeleteUserJSONRequestBody
}

type DeleteUserResponseObject interface {
	VisitDeleteUserResponse(w http.ResponseWriter) error
}

type DeleteUser204Response struct {
}

func (response DeleteUser204Response) VisitDeleteUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type GetUserRequestObject struct {
}

type GetUserResponseObject interface {
	VisitGetUserResponse(w http.ResponseWriter) error
}

type GetUser200JSONResponse []User

func (response GetUser200JSONResponse) VisitGetUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchUserRequestObject struct {
	Body *PatchUserJSONRequestBody
}

type PatchUserResponseObject interface {
	VisitPatchUserResponse(w http.ResponseWriter) error
}

type PatchUser200JSONResponse User

func (response PatchUser200JSONResponse) VisitPatchUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUserRequestObject struct {
	Body *PostUserJSONRequestBody
}

type PostUserResponseObject interface {
	VisitPostUserResponse(w http.ResponseWriter) error
}

type PostUser201JSONResponse User

func (response PostUser201JSONResponse) VisitPostUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Delete user by id
	// (DELETE /user)
	DeleteUser(ctx context.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error)
	// Get all users
	// (GET /user)
	GetUser(ctx context.Context, request GetUserRequestObject) (GetUserResponseObject, error)
	// Update user by id
	// (PATCH /user)
	PatchUser(ctx context.Context, request PatchUserRequestObject) (PatchUserResponseObject, error)
	// Create a new user
	// (POST /user)
	PostUser(ctx context.Context, request PostUserRequestObject) (PostUserResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// DeleteUser operation middleware
func (sh *strictHandler) DeleteUser(ctx echo.Context) error {
	var request DeleteUserRequestObject

	var body DeleteUserJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUser(ctx.Request().Context(), request.(DeleteUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteUserResponseObject); ok {
		return validResponse.VisitDeleteUserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUser operation middleware
func (sh *strictHandler) GetUser(ctx echo.Context) error {
	var request GetUserRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUser(ctx.Request().Context(), request.(GetUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUserResponseObject); ok {
		return validResponse.VisitGetUserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchUser operation middleware
func (sh *strictHandler) PatchUser(ctx echo.Context) error {
	var request PatchUserRequestObject

	var body PatchUserJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchUser(ctx.Request().Context(), request.(PatchUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchUserResponseObject); ok {
		return validResponse.VisitPatchUserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostUser operation middleware
func (sh *strictHandler) PostUser(ctx echo.Context) error {
	var request PostUserRequestObject

	var body PostUserJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUser(ctx.Request().Context(), request.(PostUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUserResponseObject); ok {
		return validResponse.VisitPostUserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
