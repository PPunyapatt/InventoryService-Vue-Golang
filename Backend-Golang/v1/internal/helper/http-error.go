package helper

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ErrorCodeUniqueViolation     = "23505"
	ErrorCodeForeignKeyViolation = "23503"
	ErrorCodeNotNullViolation    = "23502"
	ErrDuplicateInventory        = "23505"

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
	Message      string `json:"message"`
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
			Message:      httpError.Message,
			ErrorMessage: err.Error(),
		})
	}

	var pgError *pgconn.PgError
	ok = errors.As(err, &pgError)
	if ok {
		switch pgError.Code {
		case ErrDuplicateInventory:
			return ctx.Status(http.StatusConflict).JSON(HttpErrorResponse{
				StatusCode:   http.StatusConflict,
				Message:      fmt.Sprintf(`Duplicate Inventory Code`),
				ErrorMessage: err.Error(),
			})
		}
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return ctx.Status(http.StatusGatewayTimeout).JSON(HttpErrorResponse{
			StatusCode:   http.StatusGatewayTimeout,
			Message:      fmt.Sprintf(`Request timeout`),
			ErrorMessage: err.Error(),
		})
	}

	return ctx.Status(http.StatusInternalServerError).JSON(HttpErrorResponse{
		StatusCode:   http.StatusInternalServerError,
		Message:      ErrorMessageInternalServer,
		ErrorMessage: ErrorMessageInternalServer,
	})
}

func NewHttpError(statusCode int, message *string) HttpError {
	if message == nil {
		defaultMessage := getDefaultStatusMessage(statusCode)
		message = &defaultMessage
	}
	return HttpError{
		StatusCode: statusCode,
		Message:    *message,
	}
}

func getDefaultStatusMessage(statusCode int) string {
	message := ""
	switch statusCode {
	case http.StatusConflict:
		message = ErrorMessageConflict
	case http.StatusBadRequest:
		message = ErrorMessageBadRequest
	case http.StatusNotFound:
		message = ErrorMessageNotfound
	default:
		message = ErrorMessageInternalServer
	}

	return message
}
