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

type StructValidator struct {
	validationFuncsMap map[string]ValidationFunc
}

type ValidationFunc func(field reflect.Value, param, fieldName string) *ValidationError

// New - Returns a new validator.
func New() *StructValidator {
	stv := &StructValidator{
		validationFuncsMap: make(map[string]ValidationFunc),
	}

	stv.DefaultValidationFuncs()

	return stv
}

// AddCustomValidation - Adds a custom validation function.
func (v *StructValidator) AddCustomValidation(tag string, fn ValidationFunc) {
	v.validationFuncsMap[tag] = fn
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
				validateFunc, ok := v.validationFuncsMap[rule]
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

// DefaultValidationFuncs - Returns a map of the default validation functions.
func (v *StructValidator) DefaultValidationFuncs() {
	vf := NewValidationFuncs()

	v.validationFuncsMap["required"] = vf.IsRequired
	v.validationFuncsMap["email"] = vf.IsEmail
	v.validationFuncsMap["url"] = vf.IsURL
	v.validationFuncsMap["alpha"] = vf.IsAlpha
	v.validationFuncsMap["alphanum"] = vf.IsAlphaNumeric
	v.validationFuncsMap["numeric"] = vf.IsNumeric
	v.validationFuncsMap["hexadecimal"] = vf.IsHexadecimal
	v.validationFuncsMap["hexcolor"] = vf.IsHexColor
	v.validationFuncsMap["rgb"] = vf.IsRGBColor
	v.validationFuncsMap["rgba"] = vf.IsRGBAColor
	v.validationFuncsMap["hsl"] = vf.IsHSLColor
	v.validationFuncsMap["hsla"] = vf.IsHSLAColor
	v.validationFuncsMap["json"] = vf.IsJSON
	v.validationFuncsMap["multibyte"] = vf.IsMultibyte
	v.validationFuncsMap["ascii"] = vf.IsASCII
	v.validationFuncsMap["printableascii"] = vf.IsPrintableASCII
	v.validationFuncsMap["fullwidth"] = vf.IsFullWidth
	v.validationFuncsMap["variablewidth"] = vf.IsVariableWidth
	v.validationFuncsMap["base64"] = vf.IsBase64
	v.validationFuncsMap["datauri"] = vf.IsDataURI
	v.validationFuncsMap["ip"] = vf.IsIP
	v.validationFuncsMap["ipv4"] = vf.IsIPv4
	v.validationFuncsMap["ipv6"] = vf.IsIPv6
	v.validationFuncsMap["cidr"] = vf.IsCIDR
	v.validationFuncsMap["cidrv4"] = vf.IsCIDRv4
	v.validationFuncsMap["cidrv6"] = vf.IsCIDRv6
	v.validationFuncsMap["tcpaddr"] = vf.IsTCPAddr
	v.validationFuncsMap["udpaddr"] = vf.IsUDPAddr
	v.validationFuncsMap["latitude"] = vf.IsLatitude
	v.validationFuncsMap["longitude"] = vf.IsLongitude
	v.validationFuncsMap["ssn"] = vf.IsSSN
	v.validationFuncsMap["semver"] = vf.IsSemver
	v.validationFuncsMap["isbn10"] = vf.IsISBN10
	v.validationFuncsMap["isbn13"] = vf.IsISBN13
	v.validationFuncsMap["uuid"] = vf.IsUUID
	v.validationFuncsMap["uuid3"] = vf.IsUUID3
	v.validationFuncsMap["uuid4"] = vf.IsUUID4
	v.validationFuncsMap["uuid5"] = vf.IsUUID5
	v.validationFuncsMap["creditcard"] = vf.IsCreditCard
	v.validationFuncsMap["isbn"] = vf.IsISBN
	v.validationFuncsMap["minlen"] = vf.MinLen
	v.validationFuncsMap["maxlen"] = vf.MaxLen
	v.validationFuncsMap["len"] = vf.Len
	v.validationFuncsMap["eqfield"] = vf.EqField
	v.validationFuncsMap["eqcsfield"] = vf.EqCrossStructField
	v.validationFuncsMap["necsfield"] = vf.NeCrossStructField
	v.validationFuncsMap["gtfield"] = vf.GtField
	v.validationFuncsMap["gtefield"] = vf.GteField
	v.validationFuncsMap["ltfield"] = vf.LtField
	v.validationFuncsMap["ltefield"] = vf.LteField
	v.validationFuncsMap["nefield"] = vf.NeField
	v.validationFuncsMap["contains"] = vf.Contains
	v.validationFuncsMap["eq"] = vf.Eq
	v.validationFuncsMap["ne"] = vf.Ne
	v.validationFuncsMap["lt"] = vf.Lt
	v.validationFuncsMap["lte"] = vf.Lte
	v.validationFuncsMap["gt"] = vf.Gt
	v.validationFuncsMap["gte"] = vf.Gte
	v.validationFuncsMap["default"] = vf.Default
}
