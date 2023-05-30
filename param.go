package nex

import "github.com/google/uuid"

type param struct {
	// The name of the parameter.
	Name string `json:"name"`
	// The value of the parameter.
	Value string `json:"value"`
	// Context of the parameter.
	Context Context `json:"context"`
}

// newParam - Returns a new param.
func newParam(context Context) *param {
	return &param{
		Context: context,
	}
}

// GetParam - Returns the value of the parameter from Context.
func (p param) GetParam(key string) param {
	p.Name = key
	p.Value = p.Context.ParamGet(key)
	return p
}

// AsString - Returns the value of the parameter as a string.
func (p param) AsString() string {
	return p.Value
}

// AsInt - Returns the value of the parameter as an int.
func (p param) AsInt() (int64, error) {
	return p.Context.ParamGetInt(p.Name)
}

// AsFloat - Returns the value of the parameter as a float.
func (p param) AsBool() (bool, error) {
	return p.Context.ParamGetBool(p.Name)
}

// AsUUID - Returns the value of the parameter as a UUID.
func (p param) AsUUID() (uuid.UUID, error) {
	return p.Context.ParamGetUuid(p.Name)
}
