package service

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
)

type CreateUserValidation struct {
	Username        string `validate:"required,min=3,max=20,alphanum"`
	Email           string `validate:"required,email"`
	Age             int    `validate:"required,gte=18,lte=100"`
	Gender          string `validate:"required,oneof=M F"`
	Phone           string `validate:"required,e164"`
	Password        string `validate:"required,min=8,max=72"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
}

var validate = validator.New()

func ValidatUserRegistartion(input CreateUserValidation) error {
	input.Username = strings.TrimSpace(input.Username)
	input.Email = strings.ToLower(strings.TrimSpace(input.Email))
	input.Gender = strings.TrimSpace(input.Gender)
	input.Phone = strings.TrimSpace(input.Phone)

	if err := validate.Struct(input); err != nil {
		return formatValidationErrors(err)
	}
	if err := validatePassword(input.Password); err != nil {
		return err
	}
	return nil
}

func formatValidationErrors(err error) error {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}
	var messages []string
	for _, e := range validationErrors {
		var msg string
		switch e.Tag() {
		case "required":
			msg = fmt.Sprintf("%s is required", e.Field())
		case "email":
			msg = "invalid email format"
		case "min":
			msg = fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
		case "max":
			msg = fmt.Sprintf("%s must be at most %s characters", e.Field(), e.Param())
		case "gte":
			msg = fmt.Sprintf("%s must be at least %s", e.Field(), e.Param())
		case "lte":
			msg = fmt.Sprintf("%s must be at most %s", e.Field(), e.Param())
		case "alphanum":
			msg = fmt.Sprintf("%s must contain only letters and numbers", e.Field())
		case "oneof":
			msg = fmt.Sprintf("%s must be one of: %s", e.Field(), e.Param())
		case "e164":
			msg = "phone number must be in valid international format (e.g., +1234567890)"
		case "eqfield":
			msg = fmt.Sprintf("Passwords do not match", e.Field(), e.Param())
		default:
			msg = fmt.Sprintf("%s validation failed: %s", e.Field(), e.Tag())
		}
		messages = append(messages, msg)
	}
	return errors.New(strings.Join(messages, "; "))
}

func validatePassword(password string) error {
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)

	if !hasUpper {
		return errors.New("Password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("Password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("Passsword must contain at least one number")
	}
	if !hasSpecial {
		return errors.New("Password msut contain at least one special character")
	}
	return nil
}