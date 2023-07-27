package context

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/url"
)

type Body struct {
	ctx *Context
}

// NewBody creates a new instance of Body.
func NewBody(ctx *Context) *Body {
	return &Body{ctx: ctx}
}

// ReadRaw reads the raw request body and returns it as bytes.
func (b *Body) ReadRaw() ([]byte, error) {
	bodyBytes, err := io.ReadAll(b.ctx.Request.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

// ParseJSON parses the request body as JSON into the provided struct.
func (b *Body) ParseJSON(v interface{}) error {
	bodyBytes, err := b.ReadRaw()
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, v)
}

// ParseXML parses the request body as XML into the provided struct.
func (b *Body) ParseXML(v interface{}) error {
	bodyBytes, err := b.ReadRaw()
	if err != nil {
		return err
	}
	return xml.Unmarshal(bodyBytes, v)
}

// ParseForm parses the request body as form data.
func (b *Body) ParseForm() (values url.Values, err error) {
	err = b.ctx.Request.ParseForm()
	if err != nil {
		return nil, err
	}
	return b.ctx.Request.PostForm, nil
}
