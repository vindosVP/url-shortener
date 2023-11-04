package resp

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Response struct {
	Status string `json:"status"`
	URL    string `json:"url,omitempty"`
	Error  string `json:"errors,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OkResponse(ctx *fiber.Ctx, code int, URL string) error {
	return ctx.Status(code).JSON(Response{
		Status: StatusOK,
		URL:    URL,
	})
}

func ErrorResponse(ctx *fiber.Ctx, code int, msg string) error {
	return ctx.Status(code).JSON(Response{
		Status: StatusError,
		Error:  msg,
	})
}

func ValidationError(ctx *fiber.Ctx, code int, errs validator.ValidationErrors) error {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is a required field", err.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid URL", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return ctx.Status(code).JSON(Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	})
}
