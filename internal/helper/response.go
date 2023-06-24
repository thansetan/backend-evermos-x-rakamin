package helper

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Errors  any    `json:"errors"`
	Data    any    `json:"data"`
}

const (
	GETDATAFAILED     = "Failed to GET data"
	GETDATASUCCESS    = "Succeed to GET data"
	POSTDATAFAILED    = "Failed to POST data"
	POSTDATASUCCESS   = "Succeed to POST data"
	PUTDATASUCCESS    = "Succeed to PUT data"
	PUTDATAFAILED     = "Failed to PUT data"
	DELETEDATAFAILED  = "Failed to DELETE data"
	DELETEDATASUCCESS = "Succeed to DELETE data"
)

func ResponseBuilder(ctx fiber.Ctx, status bool, message string, err any, data any, code int) error {
	var errors []any
	if err != nil {
		errors = append(errors, err)
	}
	return ctx.Status(code).JSON(&Response{
		Status:  status,
		Message: message,
		Errors:  errors,
		Data:    data,
	})
}
