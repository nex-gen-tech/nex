package context

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"
)

type Form struct {
	ctx *Context
}

func NewForm(ctx *Context) *Form {
	return &Form{ctx: ctx}
}

// fetchFormValue fetches a form value and checks if it exists.
func (f *Form) fetchFormValue(name string) (string, error) {
	strValue := f.ctx.Request.FormValue(name)
	if strValue == "" {
		return "", fmt.Errorf("form parameter %s not found", name)
	}
	return strValue, nil
}

func (f *Form) Get(name string) string {
	return f.ctx.Request.FormValue(name)
}

// Integer methods
func (f *Form) GetAsInt(name string) (int, error) {
	strValue, err := f.fetchFormValue(name)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strValue)
}

func (f *Form) GetAsInt64(name string) (int64, error) {
	strValue, err := f.fetchFormValue(name)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(strValue, 10, 64)
}

// Boolean method
func (f *Form) GetAsBool(name string) (bool, error) {
	strValue, err := f.fetchFormValue(name)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(strValue)
}

// Bind - binds the form to the given struct.
func (f *Form) Bind(v interface{}) error {
	decoder := schema.NewDecoder()

	// Parse the form values from the request
	err := f.ctx.Request.ParseForm()
	if err != nil {
		return err
	}

	// Decode the form values into the struct
	return decoder.Decode(v, f.ctx.Request.PostForm)
}

// BindMultiPart - binds the multipart form to the given struct.
func (f *Form) BindMultiPart(v interface{}) error {
	decoder := schema.NewDecoder()

	// Parse the multipart form data from the request
	err := f.ctx.Request.ParseMultipartForm(32 << 20) // 32MB as the max memory
	if err != nil && err != http.ErrNotMultipart {
		return err
	}

	// Decode the form values into the struct
	return decoder.Decode(v, f.ctx.Request.PostForm)
}

// GetMultiple - Returns multiple values for a given key (useful for checkboxes, multi-selects).
func (f *Form) GetMultiple(name string) []string {
	return f.ctx.Request.Form[name]
}

// GetAsFloat64 - returns the value of the form parameter with the given name as a float64.
func (f *Form) GetAsFloat64(name string) (float64, error) {
	strValue, err := f.fetchFormValue(name)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(strValue, 64)
}

// GetFile - Returns the uploaded file for a given key.
func (f *Form) GetFile(name string) (multipart.File, *multipart.FileHeader, error) {
	return f.ctx.Request.FormFile(name)
}

// HasField - Checks if a field with the given name is present in the form.
func (f *Form) HasField(name string) bool {
	_, exists := f.ctx.Request.Form[name]
	return exists
}

// CheckCSRF - Checks if CSRF token is valid. This assumes you're adding a CSRF token in your forms.
func (f *Form) CheckCSRF() bool {
	// Implementation will depend on how you're handling CSRF in your application.
	// Just as a placeholder:
	csrfToken := f.Get("csrf_token")
	return csrfToken == "expected_token_value" // Replace with actual comparison logic.
}
