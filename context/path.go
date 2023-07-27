package context

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type PathParam struct {
	ctx *Context
}

func NewPathParam(ctx *Context) *PathParam {
	return &PathParam{ctx: ctx}
}

// fetchValue fetches a value from context and checks if it exists.
func (p *PathParam) fetchValue(name string) (string, error) {
	strValue := p.ctx.Params[name]
	if strValue == "" {
		return "", fmt.Errorf("path parameter %s not found", name)
	}
	return strValue, nil
}

func (p *PathParam) Get(name string) string {
	return p.ctx.Params[name]
}

// Integer related methods
func (p *PathParam) GetAsInt(name string) (int, error) {
	strValue, err := p.fetchValue(name)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strValue)
}

func (p *PathParam) GetAsInt64(name string) (int64, error) {
	strValue, err := p.fetchValue(name)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(strValue, 10, 64)
}

// Boolean method
func (p *PathParam) GetAsBool(name string) (bool, error) {
	strValue, err := p.fetchValue(name)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(strValue)
}

// UUID method
func (p *PathParam) GetAsUUID(name string) (uuid.UUID, error) {
	strValue, err := p.fetchValue(name)
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid.Parse(strValue)
}
