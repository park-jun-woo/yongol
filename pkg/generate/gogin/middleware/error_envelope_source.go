//ff:type feature=gen-gogin type=generator
//ff:what errorEnvelopeSources — error_envelope 를 6파일로 분할하는 소스 템플릿 맵

package middleware

// errorEnvelopeTypeSource — ErrorEnvelope struct + ExposeInternalError var.
const errorEnvelopeTypeSource = `//` + `ff:type feature=runtime-middleware type=model topic=error-envelope
//` + `ff:what ErrorEnvelope — 모든 실패 응답의 표준 JSON envelope 구조체

package middleware

// ExposeInternalError controls whether 500 responses carry err.Error() in
// the message field. Default false — production-safe. main.go flips it to
// true when BACKEND_ERROR_EXPOSE_INTERNAL_ERROR is set.
var ExposeInternalError = false

// ErrorEnvelope is the canonical JSON body for every failure response.
//
//   - Error     machine-readable snake_case code (e.g. "rate_limit_exceeded")
//   - Message   human-readable one-liner (default_locale, Korean by default)
//   - RequestID request id stamped by RequestID middleware (always present)
//   - RetryAfter Unix timestamp (seconds) — set on 429 only
//   - Limit     byte cap — set on 413 only
//   - FieldErrors map<field, message> — set on 422 only
type ErrorEnvelope struct {
	Error       string            ` + "`" + `json:"error"` + "`" + `
	Message     string            ` + "`" + `json:"message"` + "`" + `
	RequestID   string            ` + "`" + `json:"request_id"` + "`" + `
	RetryAfter  int64             ` + "`" + `json:"retry_after,omitempty"` + "`" + `
	Limit       int64             ` + "`" + `json:"limit,omitempty"` + "`" + `
	FieldErrors map[string]string ` + "`" + `json:"field_errors,omitempty"` + "`" + `
}
`

// defaultCodeForSource — DefaultCodeFor func.
const defaultCodeForSource = `//` + `ff:func feature=runtime-middleware type=util control=selection topic=error-envelope
//` + `ff:what DefaultCodeFor — HTTP status code 에서 machine-readable error code 로 변환

package middleware

import "net/http"

// DefaultCodeFor returns the canonical machine code for a status code.
// Falls back to "error" for unmapped statuses so the envelope always has a
// non-empty error field.
func DefaultCodeFor(status int) string {
	switch status {
	case http.StatusBadRequest:
		return "bad_request"
	case http.StatusUnauthorized:
		return "unauthorized"
	case http.StatusForbidden:
		return "forbidden"
	case http.StatusNotFound:
		return "not_found"
	case http.StatusConflict:
		return "conflict"
	case http.StatusRequestEntityTooLarge:
		return "payload_too_large"
	case http.StatusUnprocessableEntity:
		return "validation_failed"
	case http.StatusTooManyRequests:
		return "rate_limit_exceeded"
	case http.StatusInternalServerError:
		return "internal_error"
	case http.StatusServiceUnavailable:
		return "service_unavailable"
	}
	return "error"
}
`

// defaultMessageForSource — DefaultMessageFor func.
const defaultMessageForSource = `//` + `ff:func feature=runtime-middleware type=util control=selection topic=error-envelope
//` + `ff:what DefaultMessageFor — HTTP status code 에서 human-readable 메시지로 변환

package middleware

import "net/http"

// DefaultMessageFor returns the human-readable one-liner for a status code.
// Unmapped statuses yield a generic message.
func DefaultMessageFor(status int) string {
	switch status {
	case http.StatusBadRequest:
		return "Bad request"
	case http.StatusUnauthorized:
		return "Authentication required"
	case http.StatusForbidden:
		return "Access denied"
	case http.StatusNotFound:
		return "Resource not found"
	case http.StatusConflict:
		return "Resource conflict"
	case http.StatusRequestEntityTooLarge:
		return "Request body too large"
	case http.StatusUnprocessableEntity:
		return "Invalid input"
	case http.StatusTooManyRequests:
		return "Too many requests"
	case http.StatusInternalServerError:
		return "Internal server error"
	case http.StatusServiceUnavailable:
		return "Service temporarily unavailable"
	}
	return "Unable to process request"
}
`

// writeEnvelopeSource — WriteEnvelope func (thin wrapper).
const writeEnvelopeSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=error-envelope
//` + `ff:what WriteEnvelope — canonical error envelope 작성 + gin context abort

package middleware

import "github.com/gin-gonic/gin"

// WriteEnvelope writes a canonical envelope and aborts the gin context.
// code="" falls back to DefaultCodeFor(status); message="" falls back to
// DefaultMessageFor(status). request_id is sourced from the gin context
// regardless of the caller.
func WriteEnvelope(c *gin.Context, status int, code, message string) {
	WriteEnvelopeWithContext(c, status, code, message, nil)
}
`

// writeEnvelopeWithContextSource — WriteEnvelopeWithContext func.
const writeEnvelopeWithContextSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=error-envelope
//` + `ff:what WriteEnvelopeWithContext — context 필드(retry_after/limit/field_errors) 포함 canonical error envelope 작성

package middleware

import "github.com/gin-gonic/gin"

// WriteEnvelopeWithContext is the full variant — extra carries optional
// context fields (retry_after, limit, field_errors) merged into the body.
// Unknown keys are silently ignored to keep the envelope schema closed.
func WriteEnvelopeWithContext(c *gin.Context, status int, code, message string, extra map[string]interface{}) {
	if code == "" {
		code = DefaultCodeFor(status)
	}
	if message == "" {
		message = DefaultMessageFor(status)
	}
	env := ErrorEnvelope{
		Error:     code,
		Message:   message,
		RequestID: RequestIDFromContext(c),
	}
	if v, ok := extra["retry_after"]; ok {
		if n, ok := v.(int64); ok {
			env.RetryAfter = n
		}
	}
	if v, ok := extra["limit"]; ok {
		if n, ok := v.(int64); ok {
			env.Limit = n
		}
	}
	if v, ok := extra["field_errors"]; ok {
		if m, ok := v.(map[string]string); ok {
			env.FieldErrors = m
		}
	}
	c.AbortWithStatusJSON(status, env)
}
`

// errorEnvelopeMiddlewareSource — ErrorEnvelopeMiddleware func.
const errorEnvelopeMiddlewareSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=error-envelope
//` + `ff:what ErrorEnvelopeMiddleware — 미처리 abort 를 표준 JSON envelope 로 변환하는 post-handler 미들웨어

package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorEnvelopeMiddleware runs after all handlers. When c.Errors contains
// entries and nothing was written to the body, it emits a default envelope
// mapped from c.Writer.Status().
//
// Middleware position: AFTER RequestID (so RequestIDFromContext works)
// and BEFORE every other middleware that might write a 4xx/5xx body, so
// the tailing logic below catches any of them.
func ErrorEnvelopeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if !c.Writer.Written() && c.IsAborted() {
			status := c.Writer.Status()
			if status < 400 {
				status = http.StatusInternalServerError
			}
			message := DefaultMessageFor(status)
			if ExposeInternalError && len(c.Errors) > 0 {
				message = c.Errors.Last().Error()
			}
			if len(c.Errors) > 0 {
				slog.ErrorContext(c.Request.Context(), "handler error",
					"status", status, "err", c.Errors.String())
			}
			env := ErrorEnvelope{
				Error:     DefaultCodeFor(status),
				Message:   message,
				RequestID: RequestIDFromContext(c),
			}
			body, _ := json.Marshal(env)
			c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
			c.Writer.WriteHeader(status)
			_, _ = c.Writer.Write(body)
		}
	}
}
`
