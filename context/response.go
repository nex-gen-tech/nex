package context

import (
	"encoding/json"
	"net/http"
)

type NexResponse struct {
	ctx *Context
}

// NexResponse - create new NexResponse
func NewNexResponse(ctx *Context) *NexResponse {
	return &NexResponse{
		ctx: ctx,
	}
}

// JSON sends a JSON response with the given status code and payload.
func (r *NexResponse) JSON(status int, payload interface{}) {
	r.ctx.mu.Lock()
	defer r.ctx.mu.Unlock()

	r.ctx.Response.Header().Set("Content-Type", "application/json")
	r.ctx.Response.WriteHeader(status)
	if err := json.NewEncoder(r.ctx.Response).Encode(payload); err != nil {
		http.Error(r.ctx.Response, err.Error(), http.StatusInternalServerError)
	}
}

// Text sends a plain text response with the given status code and payload.
func (r *NexResponse) Text(status int, payload string) {
	r.ctx.mu.Lock()
	defer r.ctx.mu.Unlock()

	r.ctx.Response.Header().Set("Content-Type", "text/plain")
	r.ctx.Response.WriteHeader(status)
	r.ctx.Response.Write([]byte(payload))
}

// HTML sends a html response with the given status code and payload.
func (r *NexResponse) HTML(status int, payload string) {
	r.ctx.mu.Lock()
	defer r.ctx.mu.Unlock()

	r.ctx.Response.Header().Set("Content-Type", "text/html")
	r.ctx.Response.WriteHeader(status)
	r.ctx.Response.Write([]byte(payload))
}

// nexResFormat - response format
type nexResFormat struct {
	Status  int    `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

type nexResFormatGeneric[T any] struct {
	Status  int    `json:"status"`
	Data    T      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// JsonOk200 - response with status 200
// The request succeeded. The result meaning of "success" depends on the HTTP method:
// GET: The resource has been fetched and transmitted in the message body.
// HEAD: The representation headers are included in the response without any message body.
// PUT or POST: The resource describing the result of the action is transmitted in the message body.
// TRACE: The message body contains the request message as received by the server.
func (r *NexResponse) JsonOk200(data any, message string) {
	r.JSON(200, nexResFormat{
		Status:  200,
		Data:    data,
		Message: message,
	})
}

// Ok200 - response with status 200 and data of any generic type
// The request succeeded. The result meaning of "success" depends on the HTTP method:
// GET: The resource has been fetched and transmitted in the message body.
// HEAD: The representation headers are included in the response without any message body.
// PUT or POST: The resource describing the result of the action is transmitted in the message body.
// TRACE: The message body contains the request message as received by the server.
func Ok200[T any](r *NexResponse, data T, message string) {
	r.JSON(200, nexResFormatGeneric[T]{
		Status:  200,
		Data:    data,
		Message: message,
	})
}

// Created201 - response with status 201
// The request succeeded, and a new resource was created as a result.
// This is typically the response sent after POST requests, or some PUT requests.
func (r *NexResponse) JsonOk201(data any, message string) {
	r.JSON(201, nexResFormat{
		Status:  201,
		Data:    data,
		Message: message,
	})
}

// Created201 - response with status 201 and data of any generic type
// The request succeeded, and a new resource was created as a result.
// This is typically the response sent after POST requests, or some PUT requests.
func Ok201[T any](r *NexResponse, data T, message string) {
	r.JSON(201, nexResFormatGeneric[T]{
		Status:  201,
		Data:    data,
		Message: message,
	})
}

// Accepted202 - response with status 202
// The request has been received but not yet acted upon.
// It is non-committal, meaning that there is no way in HTTP to later send an asynchronous response indicating the outcome of processing the request.
// It is intended for cases where another process or server handles the request, or for batch processing.
func (r *NexResponse) JsonOk202(data any, message string) {
	r.JSON(202, nexResFormat{
		Status:  202,
		Data:    data,
		Message: message,
	})
}

// Accepted202 - response with status 202 and data of any generic type
// The request has been received but not yet acted upon.
// It is non-committal, meaning that there is no way in HTTP to later send an asynchronous response indicating the outcome of processing the request.
// It is intended for cases where another process or server handles the request, or for batch processing.
func Ok202[T any](r *NexResponse, data T, message string) {
	r.JSON(202, nexResFormatGeneric[T]{
		Status:  202,
		Data:    data,
		Message: message,
	})
}

// NoContent204 - response with status 204
// There is no content to send for this request, but the headers may be useful.
// The user-agent may update its cached headers for this resource with the new ones.
func (r *NexResponse) JsonNoContent204(message string) {
	r.JSON(204, nexResFormat{
		Status:  204,
		Message: message,
	})
}

// NoContent204 - response with status 204 and data of any generic type
// There is no content to send for this request, but the headers may be useful.
// The user-agent may update its cached headers for this resource with the new ones.
func NoContent204[T any](r *NexResponse, message string) {
	r.JSON(204, nexResFormatGeneric[T]{
		Status:  204,
		Message: message,
	})
}

// BadRequest400 - response with status 400
// The server could not understand the request due to invalid syntax.
// The client SHOULD NOT repeat the request without modifications.
func (r *NexResponse) JsonBadRequest400(message string) {
	r.JSON(400, nexResFormat{
		Status:  400,
		Message: message,
	})
}

// BadRequest400 - response with status 400 and data of any generic type
// The server could not understand the request due to invalid syntax.
// The client SHOULD NOT repeat the request without modifications.
func BadRequest400[T any](r *NexResponse, message string) {
	r.JSON(400, nexResFormatGeneric[T]{
		Status:  400,
		Message: message,
	})
}

// Unauthorized401 - response with status 401
// Although the HTTP standard specifies "unauthorized", semantically this response means "unauthenticated".
// That is, the client must authenticate itself to get the requested response.
func (r *NexResponse) JsonUnauthorized401(message string) {
	r.JSON(401, nexResFormat{
		Status:  401,
		Message: message,
	})
}

// Unauthorized401 - response with status 401 and data of any generic type
// Although the HTTP standard specifies "unauthorized", semantically this response means "unauthenticated".
// That is, the client must authenticate itself to get the requested response.
func Unauthorized401[T any](r *NexResponse, message string) {
	r.JSON(401, nexResFormatGeneric[T]{
		Status:  401,
		Message: message,
	})
}

// Forbidden403 - response with status 403
// The client does not have access rights to the content; that is, it is unauthorized, so the server is refusing to give the requested resource.
// Unlike 401, the client's identity is known to the server.
func (r *NexResponse) JsonForbidden403(message string) {
	r.JSON(403, nexResFormat{
		Status:  403,
		Message: message,
	})
}

// Forbidden403 - response with status 403 and data of any generic type
// The client does not have access rights to the content; that is, it is unauthorized, so the server is refusing to give the requested resource.
// Unlike 401, the client's identity is known to the server.
func Forbidden403[T any](r *NexResponse, message string) {
	r.JSON(403, nexResFormatGeneric[T]{
		Status:  403,
		Message: message,
	})
}

// NotFound404 - response with status 404
// The server can not find the requested resource.
// In the browser, this means the URL is not recognized.
// In an API, this can also mean that the endpoint is valid but the resource itself does not exist.
// Servers may also send this response instead of 403 to hide the existence of a resource from an unauthorized client.
// This response code is probably the most famous one due to its frequent occurrence on the web.
func (r *NexResponse) JsonNotFound404(message string) {
	r.JSON(404, nexResFormat{
		Status:  404,
		Message: message,
	})
}

// NotFound404 - response with status 404 and data of any generic type
// The server can not find the requested resource.
// In the browser, this means the URL is not recognized.
// In an API, this can also mean that the endpoint is valid but the resource itself does not exist.
// Servers may also send this response instead of 403 to hide the existence of a resource from an unauthorized client.
// This response code is probably the most famous one due to its frequent occurrence on the web.
func NotFound404[T any](r *NexResponse, message string) {
	r.JSON(404, nexResFormatGeneric[T]{
		Status:  404,
		Message: message,
	})
}

// MethodNotAllowed405 - response with status 405
// The request method is known by the server but has been disabled and cannot be used.
// For example, an API may forbid DELETE-ing a resource.
// The two mandatory methods, GET and HEAD, must never be disabled and should not return this error code.
func (r *NexResponse) JsonMethodNotAllowed405(message string) {
	r.JSON(405, nexResFormat{
		Status:  405,
		Message: message,
	})
}

// MethodNotAllowed405 - response with status 405 and data of any generic type
// The request method is known by the server but has been disabled and cannot be used.
// For example, an API may forbid DELETE-ing a resource.
// The two mandatory methods, GET and HEAD, must never be disabled and should not return this error code.
func MethodNotAllowed405[T any](r *NexResponse, message string) {
	r.JSON(405, nexResFormatGeneric[T]{
		Status:  405,
		Message: message,
	})
}

// Conflict409 - response with status 409
// This response is sent when a request conflicts with the current state of the server.
func (r *NexResponse) JsonConflict409(message string) {
	r.JSON(409, nexResFormat{
		Status:  409,
		Message: message,
	})
}

// Conflict409 - response with status 409 and data of any generic type
// This response is sent when a request conflicts with the current state of the server.
func Conflict409[T any](r *NexResponse, message string) {
	r.JSON(409, nexResFormatGeneric[T]{
		Status:  409,
		Message: message,
	})
}

// InternalServerError500 - response with status 500
// The server has encountered a situation it doesn't know how to handle.
func (r *NexResponse) JsonInternalServerError500(message string) {
	r.JSON(500, nexResFormat{
		Status:  500,
		Message: message,
	})
}

// InternalServerError500 - response with status 500 and data of any generic type
// The server has encountered a situation it doesn't know how to handle.
func InternalServerError500[T any](r *NexResponse, message string) {
	r.JSON(500, nexResFormatGeneric[T]{
		Status:  500,
		Message: message,
	})
}
