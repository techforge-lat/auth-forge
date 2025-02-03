package middlewares

import (
	"auth-forge/internal/shared/domain/ports/out"
	"errors"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/errortrace/v2/errtype"
	"github.com/techforge-lat/rapi"
	"github.com/techforge-lat/valid"
)

// HTTPErrorHandler handler the error response of echo.
func HTTPErrorHandler(logger out.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		requestID := c.Response().Header().Get(echo.HeaderXRequestID)

		var traceError *errortrace.Error
		if !errors.As(err, &traceError) {
			logger.Error(
				c.Request().Context(),
				rapi.DetailByStatusCode[string(errtype.InternalError)],
				"request_id", requestID,
				"code", string(errtype.InternalError),
				"cause", err.Error(),
			)

			response := rapi.UnexpectedError()
			if c.Echo().Debug {
				response.DebugError = err.Error()
			}

			err = c.JSON(http.StatusInternalServerError, response) //nolint:staticcheck,wastedassign,ineffassign

			return
		}

		if traceError == nil || traceError.Cause == nil {
			return
		}

		response := ToResponse(traceError)

		httpStatusCode := rapi.HTTPStatusByStatusCode[rapi.Code(response.StatusCode)]
		response.Status = uint(httpStatusCode)

		if !c.Echo().Debug {
			response.DebugError = ""
		}

		err = c.JSON(httpStatusCode, response)

		// for public routes it will not exist,
		userIDStr := c.Request().Header.Get("X-User-ID")
		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			userID = 0
		}

		attrs := LogFields(traceError)
		attrs = append(attrs, "request_id", requestID)
		attrs = append(attrs, "user_id", userID)
		attrs = append(attrs, "status_http", httpStatusCode)

		var validationErrors valid.ValidationErrors
		if errors.As(err, &validationErrors) {
			attrs = append(attrs, validationErrors.LogFields()...)
		}

		if rapi.Code(response.StatusCode) == rapi.InternalError {
			logger.Error(
				c.Request().Context(),
				response.Detail,
				attrs...,
			)

			return
		}

		logger.Warn(
			c.Request().Context(),
			response.Detail,
			attrs...,
		)
	}
}

func ToResponse(b *errortrace.Error) rapi.Response {
	if b.Code == "" {
		b.Code = string(rapi.InternalError)
	}

	if b.Title == "" {
		b.Title = rapi.TitleByStatusCode[b.Code]
	}

	if b.Message == "" {
		b.Message = rapi.DetailByStatusCode[b.Code]
	}

	response := rapi.New()

	response.Title = b.Title
	response.Detail = b.Message
	response.StatusCode = b.Code

	if errors.Is(b.Cause, pgx.ErrNoRows) {
		notFoundResponse := *rapi.New()

		notFoundResponse.Title = "Recurso no encontrado"
		notFoundResponse.Detail = "El recurso solicitado no existe"
		notFoundResponse.StatusCode = string(rapi.NotFound)

		return notFoundResponse
	}

	var validationErrors valid.ValidationErrors
	if errors.As(b.Cause, &validationErrors) {
		response.Data = validationErrors
		response.Title = "Error de validación"
		response.Detail = validationErrors.Error()
		response.StatusCode = string(rapi.UnprocessableEntity)
	}

	return *response
}

func LogFields(b *errortrace.Error) []any {
	fields := make([]interface{}, 0)

	if b.Code == "" {
		b.Code = string(errtype.InternalError)
	}

	fields = append(fields, "code", b.Code)

	if b.HasTitle() {
		fields = append(fields, "title", b.Title)
	}

	fields = append(fields, "message", b.Message)

	if b.Cause != nil {
		fields = append(fields, "cause", b.Cause.Error())
	}

	if len(b.Stack) > 0 {
		frame := b.Stack[0]
		fields = append(fields,
			"file", filepath.Base(frame.File),
			"line", frame.Line,
			"function", filepath.Base(frame.Function),
		)

		// Add full stack trace
		stackFrames := make([]map[string]interface{}, len(b.Stack))
		for i := len(b.Stack) - 1; i >= 0; i-- {
			frame := b.Stack[i]
			stackFrames[i] = map[string]interface{}{
				"file":     filepath.Base(frame.File),
				"line":     frame.Line,
				"function": filepath.Base(frame.Function),
			}
		}
		fields = append(fields, "stack", stackFrames)
	}

	return fields
}
