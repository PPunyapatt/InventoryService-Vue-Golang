package helper

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	ErrorCodeUniqueViolation     = "23505"
	ErrorCodeForeignKeyViolation = "23503"
	ErrorCodeNotNullViolation    = "23502"

	ErrorMessageConflict       = "Conflict"
	ErrorMessageBadRequest     = "Bad request"
	ErrorMessageUnauthorized   = "Unauthorized"
	ErrorMessageNotfound       = "Not found"
	ErrorMessageForbidden      = "Forbidden"
	ErrorMessageInternalServer = "Internal server error"
)

type HttpError struct {
	StatusCode   int
	Message      string
	ErrorMessage string
}

type HttpErrorResponse struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func (err HttpError) Error() string {
	return err.Message
}

func RespondHttpError(ctx *fiber.Ctx, err error) error {
	var httpError HttpError
	ok := errors.As(err, &httpError)
	if ok {
		return ctx.Status(httpError.StatusCode).JSON(HttpErrorResponse{
			StatusCode:   httpError.StatusCode,
			ErrorMessage: err.Error(),
		})
	}

	return ctx.Status(http.StatusInternalServerError).JSON(HttpErrorResponse{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: ErrorMessageInternalServer,
	})
}

func NewHttpError(statusCode int, err error) HttpError {
	return HttpError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}
