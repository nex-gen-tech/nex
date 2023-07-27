package nexval

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/spf13/cast"
)

type ValidationFuncs struct{}

// NewValidationFuncs - Returns a new validation function.
func NewValidationFuncs() *ValidationFuncs {
	return &ValidationFuncs{}
}

// IsRequired - check if the field is required
func (vf ValidationFuncs) IsRequired(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	if field.String() == "" {
		return &ValidationError{
			Field: fieldName,
			Tag:   "required",
			Err:   "Field is required",
		}
	}
	return nil
}

// IsEmail - check if the field is email with regex.
func (vf ValidationFuncs) IsEmail(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	regex := regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
	if !regex.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "email",
			Err:   "given string is not a valid email address",
		}
	}
	return nil
}

// IsURL - check if the field is a valid URL.
func (vf ValidationFuncs) IsURL(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	_, err := url.ParseRequestURI(field.String())
	if err != nil {
		return &ValidationError{
			Field: fieldName,
			Tag:   "url",
			Err:   "given string is not a valid URL",
		}
	}
	return nil
}

// IsAlpha - check if the field contains only letters (a-zA-Z).
func (vf ValidationFuncs) IsAlpha(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "alpha",
			Err:   "given string contains non-alphabetic characters",
		}
	}
	return nil
}

// IsAlphaNumeric - check if the field contains only letters and numbers.
func (vf ValidationFuncs) IsAlphaNumeric(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "alphaNum",
			Err:   "given string contains non-alphanumeric characters",
		}
	}
	return nil
}

// IsNumeric - check if the field contains only numbers.
func (vf ValidationFuncs) IsNumeric(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[0-9]+$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "numeric",
			Err:   "given string contains non-numeric characters",
		}
	}
	return nil
}

// IsHexadecimal - check if the field is a hexadecimal.
func (vf ValidationFuncs) IsHexadecimal(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[a-fA-F0-9]+$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "hexadecimal",
			Err:   "given string contains non-hexadecimal characters",
		}
	}
	return nil
}

// IsHexColor - check if the field is a hexadecimal color.
func (vf ValidationFuncs) IsHexColor(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^#?([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "hexcolor",
			Err:   "given string is not a valid hexadecimal color",
		}
	}
	return nil
}

// IsRGBColor - check if the field is a valid RGB color.
func (vf ValidationFuncs) IsRGBColor(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^rgb\((\d{1,3}),(\d{1,3}),(\d{1,3})\)$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "rgb",
			Err:   "given string is not a valid RGB color",
		}
	}
	return nil
}

// IsRGBAColor - check if the field is a valid RGBA color.
func (vf ValidationFuncs) IsRGBAColor(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^rgba\((\d{1,3}),(\d{1,3}),(\d{1,3}),(\d{1,3})\)$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "rgba",
			Err:   "given string is not a valid RGBA color",
		}
	}
	return nil
}

// IsHSLColor - check if the field is a valid HSL color.
func (vf ValidationFuncs) IsHSLColor(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^hsl\((\d{1,3}),(\d{1,3})%,(\d{1,3})%\)$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "hsl",
			Err:   "given string is not a valid HSL color",
		}
	}
	return nil
}

// IsHSLAColor - check if the field is a valid HSLA color.
func (vf ValidationFuncs) IsHSLAColor(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^hsla\((\d{1,3}),(\d{1,3})%,(\d{1,3})%,(\d{1,3})\)$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "hsla",
			Err:   "given string is not a valid HSLA color",
		}
	}
	return nil
}

// IsJSON - check if the field is a valid JSON string.
func (vf ValidationFuncs) IsJSON(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	var js json.RawMessage
	if err := json.Unmarshal([]byte(field.String()), &js); err != nil {
		return &ValidationError{
			Field: fieldName,
			Tag:   "json",
			Err:   "given string is not a valid JSON string",
		}
	}
	return nil
}

// IsMultibyte - check if the field contains one or more multibyte chars.
func (vf ValidationFuncs) IsMultibyte(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`[^\x00-\x7F]`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "multibyte",
			Err:   "given string does not contain multibyte characters",
		}
	}
	return nil
}

// IsASCII - check if the field contains ASCII chars only.
func (vf ValidationFuncs) IsASCII(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[\x00-\x7F]+$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "ascii",
			Err:   "given string does not contain ASCII characters",
		}
	}
	return nil
}

// IsPrintableASCII - check if the field contains printable ASCII chars only.
func (vf ValidationFuncs) IsPrintableASCII(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[ !\"#$%&\'()*+,\-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~]+$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "printascii",
			Err:   "given string does not contain printable ASCII characters",
		}
	}
	return nil
}

// IsFullWidth - check if the field contains any full-width chars.
func (vf ValidationFuncs) IsFullWidth(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`[^\x00-\x7F]`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "fullwidth",
			Err:   "given string does not contain full-width characters",
		}
	}
	return nil
}

// IsLowerCase - check if the field contains only lower case chars.
func (vf ValidationFuncs) IsLowerCase(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[a-z]+$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "lowercase",
			Err:   "given string does not contain lower case characters",
		}
	}
	return nil
}

// IsUpperCase - check if the field contains only upper case chars.
func (vf ValidationFuncs) IsUpperCase(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[A-Z]+$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "uppercase",
			Err:   "given string does not contain upper case characters",
		}
	}
	return nil
}

// IsVariableWidth - check if the field contains a mixture of full and half-width chars.
func (vf ValidationFuncs) IsVariableWidth(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`[^\x00-\x7F]`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "variablewidth",
			Err:   "given string does not contain a mixture of full and half-width characters",
		}
	}
	return nil
}

// IsBase64 - check if the field is a valid base64 encoded string.
func (vf ValidationFuncs) IsBase64(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[a-zA-Z0-9+/]*={0,2}$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "base64",
			Err:   "given string is not a valid base64 encoded string",
		}
	}
	return nil
}

// IsDataURI - check if the field is a valid data uri string.
func (vf ValidationFuncs) IsDataURI(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^data:[a-z]+\/[a-z]+;base64,([a-zA-Z0-9+/]*={0,2})$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "datauri",
			Err:   "given string is not a valid data uri string",
		}
	}
	return nil
}

// IsIP - check if the field is a valid IP address.
func (vf ValidationFuncs) IsIP(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	if net.ParseIP(field.String()) == nil {
		return &ValidationError{
			Field: fieldName,
			Tag:   "ip",
			Err:   "given string is not a valid IP address",
		}
	}
	return nil
}

// IsIPv4 - check if the field is a valid v4 IP address.
func (vf ValidationFuncs) IsIPv4(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	ip := net.ParseIP(field.String())
	if ip == nil || ip.To4() == nil {
		return &ValidationError{
			Field: fieldName,
			Tag:   "ipv4",
			Err:   "given string is not a valid IPv4 address",
		}
	}
	return nil
}

// IsIPv6 - check if the field is a valid v6 IP address.
func (vf ValidationFuncs) IsIPv6(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	ip := net.ParseIP(field.String())
	if ip == nil || ip.To16() == nil {
		return &ValidationError{
			Field: fieldName,
			Tag:   "ipv6",
			Err:   "given string is not a valid IPv6 address",
		}
	}
	return nil
}

// IsCIDR - check if the field is a valid CIDR address.
func (vf ValidationFuncs) IsCIDR(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	if _, _, err := net.ParseCIDR(field.String()); err != nil {
		return &ValidationError{
			Field: fieldName,
			Tag:   "cidr",
			Err:   "given string is not a valid CIDR address",
		}
	}
	return nil
}

// IsCIDRv4 - check if the field is a valid v4 CIDR address.
func (vf ValidationFuncs) IsCIDRv4(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	_, ipnet, err := net.ParseCIDR(field.String())
	if err != nil || ipnet.IP.To4() == nil {
		return &ValidationError{
			Field: fieldName,
			Tag:   "cidrv4",
			Err:   "given string is not a valid IPv4 CIDR address",
		}
	}
	return nil
}

// IsCIDRv6 - check if the field is a valid v6 CIDR address.
func (vf ValidationFuncs) IsCIDRv6(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	_, ipnet, err := net.ParseCIDR(field.String())
	if err != nil || ipnet.IP.To16() == nil {
		return &ValidationError{
			Field: fieldName,
			Tag:   "cidrv6",
			Err:   "given string is not a valid IPv6 CIDR address",
		}
	}
	return nil
}

// IsTCPAddr - check if the field is a valid TCP address.
func (vf ValidationFuncs) IsTCPAddr(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	if _, err := net.ResolveTCPAddr("tcp", field.String()); err != nil {
		return &ValidationError{
			Field: fieldName,
			Tag:   "tcpaddr",
			Err:   "given string is not a valid TCP address",
		}
	}
	return nil
}

// IsUDPAddr - check if the field is a valid UDP address.
func (vf ValidationFuncs) IsUDPAddr(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	if _, err := net.ResolveUDPAddr("udp", field.String()); err != nil {
		return &ValidationError{
			Field: fieldName,
			Tag:   "udpaddr",
			Err:   "given string is not a valid UDP address",
		}
	}
	return nil
}

// IsLatitude - check if the field is a valid latitude.
func (vf ValidationFuncs) IsLatitude(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?)$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "latitude",
			Err:   "given string is not a valid latitude",
		}
	}
	return nil
}

// IsLongitude - check if the field is a valid longitude.
func (vf ValidationFuncs) IsLongitude(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?)$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "longitude",
			Err:   "given string is not a valid longitude",
		}
	}
	return nil
}

// IsSSN - check if the field is a valid SSN.
func (vf ValidationFuncs) IsSSN(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^\d{3}-?\d{2}-?\d{4}$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "ssn",
			Err:   "given string is not a valid SSN",
		}
	}
	return nil
}

// IsSemver - check if the field is a valid semantic version.
func (vf ValidationFuncs) IsSemver(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^v?(\d+)(\.\d+)?(\.\d+)?(-[a-zA-Z0-9]+)?$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "semver",
			Err:   "given string is not a valid semantic version",
		}
	}
	return nil
}

// IsISBN10 - check if the field is a valid ISBN10.
func (vf ValidationFuncs) IsISBN10(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^(?:[0-9]{9}X|[0-9]{10})$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "isbn10",
			Err:   "given string is not a valid ISBN10",
		}
	}
	return nil
}

// IsISBN13 - check if the field is a valid ISBN13.
func (vf ValidationFuncs) IsISBN13(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^(?:[0-9]{13})$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "isbn13",
			Err:   "given string is not a valid ISBN13",
		}
	}
	return nil
}

// IsUUID - check if the field is a valid UUID.
func (vf ValidationFuncs) IsUUID(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[a-f\d]{8}(-[a-f\d]{4}){3}-[a-f\d]{12}$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "uuid",
			Err:   "given string is not a valid UUID",
		}
	}
	return nil
}

// IsUUID3 - check if the field is a valid UUID3.
func (vf ValidationFuncs) IsUUID3(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[a-f\d]{8}-([a-f\d]{4}-){3}[a-f\d]{12}$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "uuid3",
			Err:   "given string is not a valid UUID3",
		}
	}
	return nil
}

// IsUUID4 - check if the field is a valid UUID4.
func (vf ValidationFuncs) IsUUID4(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[a-f\d]{8}-([a-f\d]{4}-){3}[a-f\d]{12}$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "uuid4",
			Err:   "given string is not a valid UUID4",
		}
	}
	return nil
}

// IsUUID5 - check if the field is a valid UUID5.
func (vf ValidationFuncs) IsUUID5(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^[a-f\d]{8}-([a-f\d]{4}-){3}[a-f\d]{12}$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "uuid5",
			Err:   "given string is not a valid UUID5",
		}
	}
	return nil
}

// IsCreditCard - check if the field is a valid credit card number.
func (vf ValidationFuncs) IsCreditCard(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^((4\d{3})|(5[1-5]\d{2})|(6011))-?\d{4}-?\d{4}-?\d{4}|3[4,7]\d{13}$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "creditcard",
			Err:   "given string is not a valid credit card number",
		}
	}
	return nil
}

// IsISBN - check if the field is a valid ISBN.
func (vf ValidationFuncs) IsISBN(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	reg := regexp.MustCompile(`^(?:[0-9]{9}X|[0-9]{10}|[0-9]{13})$`)
	if !reg.MatchString(field.String()) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "isbn",
			Err:   "given string is not a valid ISBN",
		}
	}
	return nil
}

// MinLen - check if the field have a minimum length.
func (vf ValidationFuncs) MinLen(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	if len(field.String()) < cast.ToInt(param) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "minlen",
			Err:   fmt.Sprintf("given string is shorter than %s", param),
		}
	}
	return nil
}

// MaxLen - check if the field is not longer than the given length.
func (vf ValidationFuncs) MaxLen(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	if len(field.String()) > cast.ToInt(param) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "maxlen",
			Err:   fmt.Sprintf("given string is longer than %s", param),
		}
	}
	return nil
}

// Len - check if the field is equal to the given length.
func (vf ValidationFuncs) Len(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	if len(field.String()) != cast.ToInt(param) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "len",
			Err:   fmt.Sprintf("given string is not equal to %s", param),
		}
	}
	return nil
}

// EqField - check if the field is equal to the given field's value.
func (vf ValidationFuncs) EqField(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	otherFieldVal := field.FieldByName(param).String()
	// Check if the field is equal for strings and numbers both.
	fieldOneNum := cast.ToInt(field.Int())
	fieldTwoNum := cast.ToInt(otherFieldVal)
	if field.String() != otherFieldVal && fieldOneNum != fieldTwoNum {
		return &ValidationError{
			Field: fieldName,
			Tag:   "eqfield",
			Err:   fmt.Sprintf("given string is not equal to %s", param),
		}
	}
	return nil
}

// EqCrossStructField - check if the field is equal to the given field's value in the cross struct.
func (vf ValidationFuncs) EqCrossStructField(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	otherFieldVal := field.FieldByName(param).String()
	// Check if the field is equal for strings and numbers both.
	fieldOneNum := cast.ToInt(field.Int())
	fieldTwoNum := cast.ToInt(otherFieldVal)
	if field.String() != otherFieldVal && fieldOneNum != fieldTwoNum {
		return &ValidationError{
			Field: fieldName,
			Tag:   "eqcrossfield",
			Err:   fmt.Sprintf("given string is not equal to %s", param),
		}
	}
	return nil
}

// NeCrossStructField - check if the field is not equal to the given field's value in the cross struct.
func (vf ValidationFuncs) NeCrossStructField(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	otherFieldVal := field.FieldByName(param).String()
	// Check if the field is not equal for strings and numbers both.
	fieldOneNum := cast.ToInt(field.Int())
	fieldTwoNum := cast.ToInt(otherFieldVal)
	if field.String() == otherFieldVal || fieldOneNum == fieldTwoNum {
		return &ValidationError{
			Field: fieldName,
			Tag:   "necrossfield",
			Err:   fmt.Sprintf("given string is equal to %s", param),
		}
	}

	return nil
}

// NeField - check if the field is not equal to the given field's value.
func (vf ValidationFuncs) NeField(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	otherFieldVal := field.FieldByName(param).String()
	// Check if the field is not equal for strings and numbers both.
	fieldOneNum := cast.ToInt(field.Int())
	fieldTwoNum := cast.ToInt(otherFieldVal)
	if field.String() == otherFieldVal || fieldOneNum == fieldTwoNum {
		return &ValidationError{
			Field: fieldName,
			Tag:   "nefield",
			Err:   fmt.Sprintf("given string is equal to %s", param),
		}
	}

	return nil
}

// LtField - check if the field is less than the given field's value.
func (vf ValidationFuncs) LtField(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	otherFieldVal := field.FieldByName(param).String()
	// check if the field is less than for number only
	fieldOneNum := cast.ToInt(field.Int())
	fieldTwoNum := cast.ToInt(otherFieldVal)
	if fieldOneNum > fieldTwoNum {
		return &ValidationError{
			Field: fieldName,
			Tag:   "ltfield",
			Err:   fmt.Sprintf("given string is not less than %s", param),
		}
	}
	return nil
}

// LteField - check if the field is less than or equal to the given field's value.
func (vf ValidationFuncs) LteField(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	otherFieldVal := field.FieldByName(param).String()
	// check if the field is less than or equal to for number only
	fieldOneNum := cast.ToInt(field.Int())
	fieldTwoNum := cast.ToInt(otherFieldVal)
	if fieldOneNum >= fieldTwoNum {
		return &ValidationError{
			Field: fieldName,
			Tag:   "ltefield",
			Err:   fmt.Sprintf("given string is not less than or equal to %s", param),
		}
	}

	return nil
}

// GtField - check if the field is greater than the given field's value.
func (vf ValidationFuncs) GtField(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	otherFieldVal := field.FieldByName(param).String()
	// check if the field is greater than for number only
	fieldOneNum := cast.ToInt(field.Int())
	fieldTwoNum := cast.ToInt(otherFieldVal)
	if fieldOneNum < fieldTwoNum {
		return &ValidationError{
			Field: fieldName,
			Tag:   "gtfield",
			Err:   fmt.Sprintf("given string is not greater than %s", param),
		}
	}

	return nil
}

// GteField - check if the field is greater than or equal to the given field's value.
func (vf ValidationFuncs) GteField(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	otherFieldVal := field.FieldByName(param).String()
	// check if the field is greater than or equal to for number only
	fieldOneNum := cast.ToInt(field.Int())
	fieldTwoNum := cast.ToInt(otherFieldVal)
	if fieldOneNum <= fieldTwoNum {
		return &ValidationError{
			Field: fieldName,
			Tag:   "gtefield",
			Err:   fmt.Sprintf("given string is not greater than or equal to %s", param),
		}
	}

	return nil
}

// Contains - check if the field contains the given substring.
func (vf ValidationFuncs) Contains(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	if !strings.Contains(field.String(), param) {
		return &ValidationError{
			Field: fieldName,
			Tag:   "contains",
			Err:   fmt.Sprintf("given string does not contain %s", param),
		}
	}
	return nil
}

// Eq - check if the field is equal to the given value.
func (vf ValidationFuncs) Eq(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	if field.String() != param {
		return &ValidationError{
			Field: fieldName,
			Tag:   "eq",
			Err:   fmt.Sprintf("given string is not equal to %s", param),
		}
	}
	return nil
}

// Ne - check if the field is not equal to the given value.
func (vf ValidationFuncs) Ne(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	if field.String() == param {
		return &ValidationError{
			Field: fieldName,
			Tag:   "ne",
			Err:   fmt.Sprintf("given string is equal to %s", param),
		}
	}
	return nil
}

// Lt - check if the field is less than the given value.
func (vf ValidationFuncs) Lt(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	fieldNum := cast.ToInt(field.Int())
	paramNum := cast.ToInt(param)
	if fieldNum >= paramNum {
		return &ValidationError{
			Field: fieldName,
			Tag:   "lt",
			Err:   fmt.Sprintf("given string is not less than %s", param),
		}
	}
	return nil
}

// Lte - check if the field is less than or equal to the given value.
func (vf ValidationFuncs) Lte(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	fieldNum := cast.ToInt(field.Int())
	paramNum := cast.ToInt(param)
	if fieldNum > paramNum {
		return &ValidationError{
			Field: fieldName,
			Tag:   "lte",
			Err:   fmt.Sprintf("given string is not less than or equal to %s", param),
		}
	}
	return nil
}

// Gt - check if the field is greater than the given value.
func (vf ValidationFuncs) Gt(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	fieldNum := cast.ToInt(field.Int())
	paramNum := cast.ToInt(param)
	log.Println(fieldNum, paramNum)
	if fieldNum <= paramNum {
		return &ValidationError{
			Field: fieldName,
			Tag:   "gt",
			Err:   fmt.Sprintf("given string is not greater than %s", param),
		}
	}
	return nil
}

// Gte - check if the field is greater than or equal to the given value.
func (vf ValidationFuncs) Gte(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	fieldNum := cast.ToInt(field.Int())
	paramNum := cast.ToInt(param)
	if fieldNum < paramNum {
		return &ValidationError{
			Field: fieldName,
			Tag:   "gte",
			Err:   fmt.Sprintf("given string is not greater than or equal to %s", param),
		}
	}
	return nil
}

// Default - set the default value for the field.
func (vf ValidationFuncs) Default(
	field reflect.Value, param, fieldName string,
) *ValidationError {
	if field.String() == "" {
		field.SetString(param)
	}
	return nil
}
