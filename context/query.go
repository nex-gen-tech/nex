package context

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type QueryParam struct {
	ctx *Context
}

// NewQueryParam creates a new instance of QueryParam.
func NewQueryParam(ctx *Context) *QueryParam {
	return &QueryParam{ctx: ctx}
}

// fetchQueryParam fetches a query parameter and checks if it exists.
func (q *QueryParam) fetchQueryParam(name string) (string, error) {
	strValue := q.ctx.Request.URL.Query().Get(name)
	if strValue == "" {
		return "", fmt.Errorf("query parameter %s not found", name)
	}
	return strValue, nil
}

// Get returns the value of the query parameter with the given name.
func (q *QueryParam) Get(name string) string {
	return q.ctx.Request.URL.Query().Get(name)
}

// Integer related methods
func (q *QueryParam) GetAsInt(name string) (int, error) {
	strValue, err := q.fetchQueryParam(name)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strValue)
}

func (q *QueryParam) GetAsInt64(name string) (int64, error) {
	strValue, err := q.fetchQueryParam(name)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(strValue, 10, 64)
}

// Boolean method
func (q *QueryParam) GetAsBool(name string) (bool, error) {
	strValue, err := q.fetchQueryParam(name)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(strValue)
}

// UUID method
func (q *QueryParam) GetAsUUID(name string) (uuid.UUID, error) {
	strValue, err := q.fetchQueryParam(name)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.Parse(strValue)
}
