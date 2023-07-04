package nexval

import (
	"reflect"
	"strings"
	"sync"
)

type ValidationError struct {
	Field string
	Tag   string
	Err   string
}

type Validation interface {
	Validate(v interface{}) []ValidationError
}

type StructValidator struct{}

type ValidationFunc func(field reflect.Value, param, fieldName string) *ValidationError

// New - Returns a new validator.
func New() *StructValidator {
	return &StructValidator{}
}

// AddCustomValidation - Adds a custom validation function.
func (v *StructValidator) AddCustomValidation(tag string, fn ValidationFunc) {
	validationFuncsMap[tag] = fn
}

// Validate - Validates the struct.
func (v *StructValidator) Validate(s interface{}) []ValidationError {
	var errors []ValidationError

	// Use reflection to iterate over the fields of the struct
	value := reflect.ValueOf(s)

	// Check if it's a pointer and get the underlying value
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// Create a channel to receive validation errors
	errChan := make(chan ValidationError)

	// Iterate over the fields of the struct in parallel
	var wg sync.WaitGroup
	for i := 0; i < value.NumField(); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			field := value.Field(i)
			tag := value.Type().Field(i).Tag.Get("nex")

			// Skip if there's no nexVal tag
			if tag == "" {
				return
			}

			// Split the tag into parts
			parts := strings.Split(tag, ",")
			for _, part := range parts {
				rule := part
				param := ""

				if strings.Contains(part, "=") {
					ruleParam := strings.Split(part, "=")
					rule = ruleParam[0]
					param = ruleParam[1]
				}
				validateFunc, ok := validationFuncsMap[rule]
				if !ok {
					continue
				}

				err := validateFunc(field, param, value.Type().Field(i).Name)
				if err != nil {
					errChan <- *err
				}
			}
		}(i)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Collect validation errors from the channel
	for err := range errChan {
		errors = append(errors, err)
	}

	return errors
}
