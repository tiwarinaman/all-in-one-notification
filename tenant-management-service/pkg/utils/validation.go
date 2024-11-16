package utils

import (
	"fmt"
	"regexp"
)

// ValidationError is a custom error type to capture validation error.
type ValidationError struct {
	Field   string
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Message)
}

// ValidateEmail validates an email address using a regular expression.
func ValidateEmail(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	if !match {
		return &ValidationError{Field: "Email", Message: "Invalid email format"}
	}
	return nil
}

// ValidatePhone validates a phone number. It checks if the number contains only digits and has a length of 10-15.
func ValidatePhone(phone string) error {
	phoneRegex := `^[0-9]{10,15}$`
	match, _ := regexp.MatchString(phoneRegex, phone)
	if !match {
		return &ValidationError{Field: "Phone", Message: "Invalid phone number format (must be 10-15 digits)"}
	}
	return nil
}

// ValidateNonEmptyString validates that a string is not empty.
func ValidateNonEmptyString(value string, fieldName string) error {
	if len(value) == 0 {
		return &ValidationError{Field: fieldName, Message: "Field cannot be empty"}
	}
	return nil
}

// ValidateAllowedValues validates that a string is within a set of allowed values.
func ValidateAllowedValues(value string, fieldName string, allowedValues []string) error {
	for _, allowed := range allowedValues {
		if value == allowed {
			return nil
		}
	}
	return &ValidationError{
		Field:   fieldName,
		Message: fmt.Sprintf("Field must be one of: %v", allowedValues),
	}
}

// ValidateMaxLength checks if a string does not exceed the maximum length.
func ValidateMaxLength(value string, fieldName string, maxLength int) error {
	if len(value) > maxLength {
		return &ValidationError{
			Field:   fieldName,
			Message: fmt.Sprintf("Field must not exceed %d characters", maxLength),
		}
	}
	return nil
}
