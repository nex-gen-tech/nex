# Nexval - A Struct Validation Package

The Nexval package provides a comprehensive set of utilities for validating structs in Go. These utilities enable you to specify constraints on your data structures and then verify that the data adheres to those constraints.

## Installation

To install the nexval package, use the following command:

```bash
go get github.com/yourusername/nexval
```

Replace `yourusername` with your actual GitHub username.

## Overview

The main validation function is `StructValidator.Validate(s interface{})`, which takes a struct `s` and returns a slice of `ValidationError` if any fields in the struct fail validation.

The `ValidationError` struct has three fields:

- `Field`: The name of the field that failed validation.
- `Tag`: The validation rule that failed.
- `Err`: The error message.

Validation rules are specified using struct tags with the key `nexVal`. The tag format is `nexVal:"rule1,rule2,..."`. Multiple rules can be specified, separated by commas.

Here is an example struct with validation tags:

```go
type User struct {
    Name  string `nexVal:"required,alpha"`
    Email string `nexVal:"required,email"`
    Age   int    `nexVal:"min=18"`
}
```

## Validation Rules

The following validation rules are supported:

- `required`: The field must not be empty.
- `email`: The field must be a valid email address.
- `url`: The field must be a valid URL.
- `alpha`: The field must contain only letters.
- `alphaNumeric`: The field must contain only letters and numbers.
- `numeric`: The field must contain only numbers.
- `lowerCase`: The field must contain only lowercase letters.
- `upperCase`: The field must contain only uppercase letters.
- `json`: The field must be a valid JSON.
- `base64`: The field must be a valid base64.
- `jwt`: The field must be a valid JWT.
- `uuid`: The field must be a valid UUID.
- `creditCard`: The field must be a valid credit card number.
- `isbn`: The field must be a valid ISBN.
- `dataURI`: The field must be a valid data URI.
- `macAddress`: The field must be a valid MAC address.
- `latitude`: The field must be a valid latitude.
- `longitude`: The field must be a valid longitude.
- `ssn`: The field must be a valid SSN.
- `minLen`: The field must be at least a certain length.
- `maxLen`: The field must be at most a certain length.
- `eqField`: The field must be equal to another field.
- `eqcsfield`: The field must be equal to another field (case sensitive).
- `neField`: The field must not be equal to another field.
- `gtField`: The field must be greater than another field.
- `gteField`: The field must be greater than or equal to another field.
- `ltField`: The field must be less than another field.
- `lteField`: The field must be less than or equal to another field.
- `contains`: The field must contain a certain substring.
- `containsany`: The field must contain any character of a certain string.
- `startsWith`: The field must start with a certain substring.
- `endsWith`: The field must end with a certain substring.
- `excludes`: The field must not contain a certain substring.
- `eq`: The field must be equal to a certain value.
- `ne`: The field must not be equal to a certain value.
- `gt`: The field must be greater than a certain value.
- `gte`: The field must be greater than or equal to a certain value.
- `lt`: The field must be less than a certain value.
- `lte`: The field must be less than or equal to a certain value.
- `default`: we can set default value for the field if it is empty.

```go
type User struct {
    Age int `nexVal:"gt=21"`
}

user := User{Age: 18}
errors := validator.Validate(user)
```
In this example, since the `Age` of the user is not greater than `21`, it will return an error message.

gte: The field must be greater than or equal to a certain value.

```go
type User struct {
    Age int `nexVal:"gte=21"`
}

user := User{Age: 21}
errors := validator.Validate(user)
```
In this example, since the `Age` of the user is equal to `21`, it will not return any errors.

lt: The field must be less than a certain value.

```go
type User struct {
    Age int `nexVal:"lt=21"`
}

user := User{Age: 22}
errors := validator.Validate(user)
```
In this example, since the `Age` of the user is not less than `21`, it will return an error message.

lte: The field must be less than or equal to a certain value.

```go
type User struct {
    Age int `nexVal:"lte=21"`
}

user := User{Age: 21}
errors := validator.Validate(user)
```
In this example, since the `Age` of the user is equal to `21`, it will not return any errors.

## Examples of struct validation
Let's consider a complex struct that requires various types of validation:

```go
type User struct {
    Name            string `nexVal:"required,alpha"`
    Email           string `nexVal:"required,email"`
    Password        string `nexVal:"required,minLen=8"`
    ConfirmPassword string `nexVal:"required,eqField=Password"`
    Age             int    `nexVal:"gte=18"`
    Website         string `nexVal:"url"`
}

user := User{
    Name:            "John Doe",
    Email:           "john.doe",
    Password:        "password",
    ConfirmPassword: "password1",
    Age:             17,
    Website:         "invalid_website",
}

errors := validator.Validate(user)
```
In this example, the `Name` field is correctly filled, but the `Email` is not valid, `Password` and `ConfirmPassword` do not match, `Age` is less than `18`, and `Website` is not a valid URL. Hence, the `Validate` function will return a list of `ValidationError`s.