package nex

// import (
// 	"encoding/json"
// 	"encoding/xml"
// 	"mime/multipart"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// 	"time"

// 	"github.com/google/uuid"

// 	"github.com/nex-gen-tech/nex/pkg/nexval"
// )

// type ResponseWriteNex struct {
// 	http.ResponseWriter
// 	Status int
// }

// func (w *ResponseWriteNex) WriteHeader(status int) {
// 	w.Status = status
// 	w.ResponseWriter.WriteHeader(status)
// }

// // Context is the context of the current HTTP request. It holds request and
// // response objects, path parameters, data, and other information.
// type Context struct {
// 	Req    *http.Request
// 	Res    ResponseWriteNex
// 	Params map[string]string
// 	Data   map[string]interface{}
// 	Param  *param
// 	Query  *query
// }

// // New creates a new Context object for a given HTTP request.
// func NewContext(r *http.Request, w ResponseWriteNex, params map[string]string) *Context {
// 	ctx := Context{
// 		Req:    r,
// 		Res:    w,
// 		Params: params,
// 		Data:   make(map[string]interface{}),
// 	}

// 	ctx.Param = newParam(ctx)
// 	ctx.Query = newQuery(ctx)

// 	return &ctx
// }

// // Get returns the value of a data item.
// func (c *Context) Get(key string) interface{} {
// 	return c.Data[key]
// }

// // Set sets the value of a data item.
// func (c *Context) Set(key string, value interface{}) {
// 	c.Data[key] = value
// }

// // ResponseJson - writes a JSON response to the client.
// func (c *Context) ResponseJson(status int, data interface{}) {
// 	c.Res.Header().Set("Content-Type", "application/json")
// 	c.Res.WriteHeader(status)
// 	if err := json.NewEncoder(c.Res).Encode(data); err != nil {
// 		http.Error(&c.Res, err.Error(), http.StatusInternalServerError)
// 	}
// }

// // ResponseJsonOk200 - writes a JSON response with status 200 to the client.
// func (c *Context) ResponseJsonOk200(data interface{}) {
// 	c.ResponseJson(http.StatusOK, data)
// }

// // ResponseJsonCreated201 - writes a JSON response with status 201 to the client.
// func (c *Context) ResponseJsonCreated201(data interface{}) {
// 	c.ResponseJson(http.StatusCreated, data)
// }

// // ResponseJsonBadRequest400 - writes a JSON response with status 400 to the client.
// func (c *Context) ResponseJsonBadRequest400(data interface{}) {
// 	c.ResponseJson(http.StatusBadRequest, data)
// }

// // ResponseJsonUnauthorized401 - writes a JSON response with status 401 to the client.
// func (c *Context) ResponseJsonUnauthorized401(data interface{}) {
// 	c.ResponseJson(http.StatusUnauthorized, data)
// }

// // ResponseJsonForbidden403 - writes a JSON response with status 403 to the client.
// func (c *Context) ResponseJsonForbidden403(data interface{}) {
// 	c.ResponseJson(http.StatusForbidden, data)
// }

// // ResponseJsonNotFound404 - writes a JSON response with status 404 to the client.
// func (c *Context) ResponseJsonNotFound404(data interface{}) {
// 	c.ResponseJson(http.StatusNotFound, data)
// }

// // NMap - is a map[string]interface{}.
// type MapNex map[string]interface{}

// // NexRes - is Struct for Response
// type ResNex struct {
// 	Code    int         `json:"code"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

// // ResNexOk200 - writes a JSON response with status 200 to the client.
// // It takes a data interface{} and message string as parameters.
// func (c *Context) ResNexOk200(data interface{}, message string) {
// 	c.ResponseJson(http.StatusOK, ResNex{
// 		Code:    http.StatusOK,
// 		Message: message,
// 		Data:    data,
// 	})
// }

// // ResNexCreated201 - writes a JSON response with status 201 to the client.
// // It takes a data interface{} and message string as parameters.
// func (c *Context) ResNexCreated201(data interface{}, message string) {
// 	c.ResponseJson(http.StatusCreated, ResNex{
// 		Code:    http.StatusCreated,
// 		Message: message,
// 		Data:    data,
// 	})
// }

// // ResNexBadRequest400 - writes a JSON response with status 400 to the client.
// // It takes only message string as parameters and data is nil.
// func (c *Context) ResNexBadRequest400(message string, data any) {
// 	c.ResponseJson(http.StatusBadRequest, ResNex{
// 		Code:    http.StatusBadRequest,
// 		Message: message,
// 		Data:    data,
// 	})
// }

// // ResNexUnauthorized401 - writes a JSON response with status 401 to the client.
// // It takes only message string as parameters and data is nil.
// func (c *Context) ResNexUnauthorized401(message string, data any) {
// 	c.ResponseJson(http.StatusUnauthorized, ResNex{
// 		Code:    http.StatusUnauthorized,
// 		Message: message,
// 		Data:    data,
// 	})
// }

// // ResNexForbidden403 - writes a JSON response with status 403 to the client.
// // It takes only message string as parameters and data is nil.
// func (c *Context) ResNexForbidden403(message string) {
// 	c.ResponseJson(http.StatusForbidden, ResNex{
// 		Code:    http.StatusForbidden,
// 		Message: message,
// 		Data:    nil,
// 	})
// }

// // ResNexNotFound404 - writes a JSON response with status 404 to the client.
// // It takes only message string as parameters and data is nil.
// func (c *Context) ResNexNotFound404(message string) {
// 	c.ResponseJson(http.StatusNotFound, ResNex{
// 		Code:    http.StatusNotFound,
// 		Message: message,
// 		Data:    nil,
// 	})
// }

// // ResNexInternalServerError500 - writes a JSON response with status 500 to the client.
// // It takes only message string as parameters and data is nil.
// func (c *Context) ResNexInternalServerError500(message string) {
// 	c.ResponseJson(http.StatusInternalServerError, ResNex{
// 		Code:    http.StatusInternalServerError,
// 		Message: message,
// 		Data:    nil,
// 	})
// }

// // ResponseJsonInternalServerError500 - writes a JSON response with status 500 to the client.
// func (c *Context) ResponseJsonInternalServerError500(data interface{}) {
// 	c.ResponseJson(http.StatusInternalServerError, data)
// }

// // ResponseHtml - writes a HTML response to the client.
// func (c *Context) ResponseHtml(status int, data interface{}) {
// 	c.Res.Header().Set("Content-Type", "text/html")
// 	c.Res.WriteHeader(status)
// 	c.Res.Write(data.([]byte))
// }

// // ResponseString - writes a string response to the client.
// func (c *Context) ResponseString(status int, data interface{}) {
// 	c.Res.WriteHeader(status)
// 	c.Res.Write(data.([]byte))
// }

// // Params

// // ParamGet - returns the value of a path parameter.
// func (c *Context) ParamGet(key string) string {
// 	return c.Params[key]
// }

// // ParamGetInt -  returns the value of a path parameter as an integer.
// func (c *Context) ParamGetInt(key string) (int64, error) {
// 	return strconv.ParseInt(c.Params[key], 10, 64)
// }

// // ParamMustGetInt -  returns the value of a path parameter as an integer. If the
// // value cannot be parsed as an integer, it panics.
// func (c *Context) ParamMustGetInt(key string) int64 {
// 	value, err := c.ParamGetInt(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // ParamGetUuid -  returns the value of a path parameter as a UUID.
// func (c *Context) ParamGetUuid(key string) (uuid.UUID, error) {
// 	return uuid.Parse(c.Params[key])
// }

// // ParamMustGetUuid -  returns the value of a path parameter as a UUID. If the
// // value cannot be parsed as a UUID, it panics.
// func (c *Context) ParamMustGetUuid(key string) uuid.UUID {
// 	value, err := c.ParamGetUuid(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // ParamGetBool -  returns the value of a path parameter as a boolean.
// func (c *Context) ParamGetBool(key string) (bool, error) {
// 	return strconv.ParseBool(c.Params[key])
// }

// // ParamMustGetBool -  returns the value of a path parameter as a boolean. If the
// // value cannot be parsed as a boolean, it panics.
// func (c *Context) ParamMustGetBool(key string) bool {
// 	value, err := c.ParamGetBool(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // Query

// // QueryGet - returns the value of a query parameter.
// func (c *Context) QueryGet(key string) string {
// 	return c.Req.URL.Query().Get(key)
// }

// // QueryGetInt -  returns the value of a query parameter as an integer.
// func (c *Context) QueryGetInt(key string) (int64, error) {
// 	return strconv.ParseInt(c.QueryGet(key), 10, 64)
// }

// // QueryMustGetInt -  returns the value of a query parameter as an integer. If the
// // value cannot be parsed as an integer, it panics.
// func (c *Context) QueryMustGetInt(key string) int64 {
// 	value, err := c.QueryGetInt(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // QueryGetUuid -  returns the value of a query parameter as a UUID
// func (c *Context) QueryGetUuid(key string) (uuid.UUID, error) {
// 	return uuid.Parse(c.QueryGet(key))
// }

// // QueryMustGetUuid -  returns the value of a query parameter as a UUID. If the
// // value cannot be parsed as a UUID, it panics.
// func (c *Context) QueryMustGetUuid(key string) uuid.UUID {
// 	value, err := c.QueryGetUuid(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // QueryGetBool -  returns the value of a query parameter as a boolean.
// func (c *Context) QueryGetBool(key string) (bool, error) {
// 	return strconv.ParseBool(c.QueryGet(key))
// }

// // QueryMustGetBool -  returns the value of a query parameter as a boolean. If the
// // value cannot be parsed as a boolean, it panics.
// func (c *Context) QueryMustGetBool(key string) bool {
// 	value, err := c.QueryGetBool(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // Header

// // HeaderGet - returns the value of a header.
// func (c *Context) HeaderGet(key string) string {
// 	return c.Req.Header.Get(key)
// }

// // HeaderGetInt -  returns the value of a header as an integer.
// func (c *Context) HeaderGetInt(key string) (int64, error) {
// 	return strconv.ParseInt(c.HeaderGet(key), 10, 64)
// }

// // HeaderMustGetInt -  returns the value of a header as an integer. If the
// // value cannot be parsed as an integer, it panics.
// func (c *Context) HeaderMustGetInt(key string) int64 {
// 	value, err := c.HeaderGetInt(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // HeaderGetUuid -  returns the value of a header as a UUID
// func (c *Context) HeaderGetUuid(key string) (uuid.UUID, error) {
// 	return uuid.Parse(c.HeaderGet(key))
// }

// // HeaderMustGetUuid -  returns the value of a header as a UUID. If the
// // value cannot be parsed as a UUID, it panics.
// func (c *Context) HeaderMustGetUuid(key string) uuid.UUID {
// 	value, err := c.HeaderGetUuid(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // HeaderGetBool -  returns the value of a header as a boolean.
// func (c *Context) HeaderGetBool(key string) (bool, error) {
// 	return strconv.ParseBool(c.HeaderGet(key))
// }

// // HeaderMustGetBool -  returns the value of a header as a boolean. If the
// // value cannot be parsed as a boolean, it panics.
// func (c *Context) HeaderMustGetBool(key string) bool {
// 	value, err := c.HeaderGetBool(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // HederSet - sets the value of a header.
// func (c *Context) HeaderSet(key string, value string) {
// 	c.Res.Header().Set(key, value)
// }

// // Cookie

// // CookieGet - returns the value of a cookie.
// func (c *Context) CookieGet(key string) string {
// 	cookie, err := c.Req.Cookie(key)
// 	if err != nil {
// 		return ""
// 	}
// 	return cookie.Value
// }

// // CookieGetInt -  returns the value of a cookie as an integer.
// func (c *Context) CookieGetInt(key string) (int64, error) {
// 	return strconv.ParseInt(c.CookieGet(key), 10, 64)
// }

// // CookieMustGetInt -  returns the value of a cookie as an integer. If the
// // value cannot be parsed as an integer, it panics.
// func (c *Context) CookieMustGetInt(key string) int64 {
// 	value, err := c.CookieGetInt(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // CookieGetUuid -  returns the value of a cookie as a UUID
// func (c *Context) CookieGetUuid(key string) (uuid.UUID, error) {
// 	return uuid.Parse(c.CookieGet(key))
// }

// // CookieMustGetUuid -  returns the value of a cookie as a UUID. If the
// // value cannot be parsed as a UUID, it panics.
// func (c *Context) CookieMustGetUuid(key string) uuid.UUID {
// 	value, err := c.CookieGetUuid(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // CookieGetBool -  returns the value of a cookie as a boolean.
// func (c *Context) CookieGetBool(key string) (bool, error) {
// 	return strconv.ParseBool(c.CookieGet(key))
// }

// // CookieMustGetBool -  returns the value of a cookie as a boolean. If the
// // value cannot be parsed as a boolean, it panics.
// func (c *Context) CookieMustGetBool(key string) bool {
// 	value, err := c.CookieGetBool(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // CookieSet - sets the value of a cookie.
// func (c *Context) CookieSet(key string, value string) {
// 	http.SetCookie(&c.Res, &http.Cookie{
// 		Name:  key,
// 		Value: value,
// 	})
// }

// // CookieSetInt - sets the value of a cookie as an integer.
// func (c *Context) CookieSetInt(key string, value int64) {
// 	c.CookieSet(key, strconv.FormatInt(value, 10))
// }

// // CookieSetUuid - sets the value of a cookie as a UUID.
// func (c *Context) CookieSetUuid(key string, value uuid.UUID) {
// 	c.CookieSet(key, value.String())
// }

// // CookieSetBool - sets the value of a cookie as a boolean.
// func (c *Context) CookieSetBool(key string, value bool) {
// 	c.CookieSet(key, strconv.FormatBool(value))
// }

// // CookieDelete - deletes a cookie.
// func (c *Context) CookieDelete(key string) {
// 	http.SetCookie(&c.Res, &http.Cookie{
// 		Name:    key,
// 		Value:   "",
// 		Expires: time.Unix(0, 0),
// 	})
// }

// // Session

// // SessionGet - returns the value of a session variable.
// func (c *Context) SessionGet(key string) any {
// 	return c.Data[key]
// }

// // SessionGetInt -  returns the value of a session variable as an integer.
// func (c *Context) SessionGetInt(key string) (int64, error) {
// 	return strconv.ParseInt(c.SessionGet(key).(string), 10, 64)
// }

// // SessionMustGetInt -  returns the value of a session variable as an integer. If the
// // value cannot be parsed as an integer, it panics.
// func (c *Context) SessionMustGetInt(key string) int64 {
// 	value, err := c.SessionGetInt(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // SessionGetUuid -  returns the value of a session variable as a UUID
// func (c *Context) SessionGetUuid(key string) (uuid.UUID, error) {
// 	return uuid.Parse(c.SessionGet(key).(string))
// }

// // SessionMustGetUuid -  returns the value of a session variable as a UUID. If the
// // value cannot be parsed as a UUID, it panics.
// func (c *Context) SessionMustGetUuid(key string) uuid.UUID {
// 	value, err := c.SessionGetUuid(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // SessionGetBool -  returns the value of a session variable as a boolean.
// func (c *Context) SessionGetBool(key string) (bool, error) {
// 	return strconv.ParseBool(c.SessionGet(key).(string))
// }

// // SessionMustGetBool -  returns the value of a session variable as a boolean. If the
// // value cannot be parsed as a boolean, it panics.
// func (c *Context) SessionMustGetBool(key string) bool {
// 	value, err := c.SessionGetBool(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // SessionSet - sets the value of a session variable.
// func (c *Context) SessionSet(key string, value any) {
// 	c.Data[key] = value
// }

// // SessionSetInt - sets the value of a session variable as an integer.
// func (c *Context) SessionSetInt(key string, value int64) {
// 	c.SessionSet(key, strconv.FormatInt(value, 10))
// }

// // SessionSetUuid - sets the value of a session variable as a UUID.
// func (c *Context) SessionSetUuid(key string, value uuid.UUID) {
// 	c.SessionSet(key, value.String())
// }

// // SessionSetBool - sets the value of a session variable as a boolean.
// func (c *Context) SessionSetBool(key string, value bool) {
// 	c.SessionSet(key, strconv.FormatBool(value))
// }

// // SessionDelete - deletes a session variable.
// func (c *Context) SessionDelete(key string) {
// 	delete(c.Data, key)
// }

// // SessionClear - clears all session variables.
// func (c *Context) SessionClear() {
// 	c.Data = map[string]any{}
// }

// // Form

// // FormGet - returns the value of a form field.
// func (c *Context) FormGet(key string) string {
// 	return c.Req.FormValue(key)
// }

// // FormGetInt -  returns the value of a form field as an integer.
// func (c *Context) FormGetInt(key string) (int64, error) {
// 	return strconv.ParseInt(c.FormGet(key), 10, 64)
// }

// // FormMustGetInt -  returns the value of a form field as an integer. If the
// // value cannot be parsed as an integer, it panics.
// func (c *Context) FormMustGetInt(key string) int64 {
// 	value, err := c.FormGetInt(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // FormGetUuid -  returns the value of a form field as a UUID
// func (c *Context) FormGetUuid(key string) (uuid.UUID, error) {
// 	return uuid.Parse(c.FormGet(key))
// }

// // FormMustGetUuid -  returns the value of a form field as a UUID. If the
// // value cannot be parsed as a UUID, it panics.
// func (c *Context) FormMustGetUuid(key string) uuid.UUID {
// 	value, err := c.FormGetUuid(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // FormGetBool -  returns the value of a form field as a boolean.
// func (c *Context) FormGetBool(key string) (bool, error) {
// 	return strconv.ParseBool(c.FormGet(key))
// }

// // FormMustGetBool -  returns the value of a form field as a boolean. If the
// // value cannot be parsed as a boolean, it panics.
// func (c *Context) FormMustGetBool(key string) bool {
// 	value, err := c.FormGetBool(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return value
// }

// // FormGetFile - returns the file associated with a form field.
// func (c *Context) FormGetFile(key string) (*multipart.FileHeader, error) {
// 	_, fh, err := c.Req.FormFile(key)
// 	return fh, err
// }

// // FormMustGetFile - returns the file associated with a form field. If the
// // file cannot be found, it panics.
// func (c *Context) FormMustGetFile(key string) *multipart.FileHeader {
// 	fh, err := c.FormGetFile(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return fh
// }

// // FormSet - sets the value of a form field.
// func (c *Context) FormSet(key string, value string) {
// 	c.Req.Form.Set(key, value)
// }

// // FormSetInt - sets the value of a form field as an integer.
// func (c *Context) FormSetInt(key string, value int64) {
// 	c.FormSet(key, strconv.FormatInt(value, 10))
// }

// // FormSetUuid - sets the value of a form field as a UUID.
// func (c *Context) FormSetUuid(key string, value uuid.UUID) {
// 	c.FormSet(key, value.String())
// }

// // FormSetBool - sets the value of a form field as a boolean.
// func (c *Context) FormSetBool(key string, value bool) {
// 	c.FormSet(key, strconv.FormatBool(value))
// }

// // FormDelete - deletes a form field.
// func (c *Context) FormDelete(key string) {
// 	c.Req.Form.Del(key)
// }

// // FormClear - clears all form fields.
// func (c *Context) FormClear() {
// 	c.Req.Form = url.Values{}
// }

// // Body

// // BodyBindJson - binds the body of the request to a struct.
// func (c *Context) BodyBindJson(v interface{}) error {
// 	return json.NewDecoder(c.Req.Body).Decode(v)
// }

// // BodyMustBindJson - binds the body of the request to a struct. If the body
// func (c *Context) BodyMustBindJson(v interface{}) {
// 	if err := c.BodyBindJson(v); err != nil {
// 		panic(err)
// 	}
// }

// // BodyBindXml  - binds the body of the request to a struct.
// func (c *Context) BodyBindXml(v interface{}) error {
// 	err := xml.NewDecoder(c.Req.Body).Decode(v)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // BodyMustBindXml - binds the body of the request to a struct. If the body
// func (c *Context) BodyMustBindXml(v interface{}) {
// 	if err := c.BodyBindXml(v); err != nil {
// 		panic(err)
// 	}
// }

// // Validate - validates a struct.with the Go Playground validator v10 package struct tags.
// func (c *Context) Validate(v interface{}) []nexval.ValidationError {
// 	return nexval.New().Validate(v)
// }
