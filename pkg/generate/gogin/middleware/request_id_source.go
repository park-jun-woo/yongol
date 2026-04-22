//ff:type feature=gen-gogin type=generator
//ff:what requestIDSource — GenerateRequestID 가 기록하는 request_id.go 템플릿 상수

package middleware

// requestIDSource is the verbatim Go source written to
// internal/middleware/request_id.go. Provides:
//
//   - RequestID(trustUpstream bool, header string) — gin.HandlerFunc that
//     stamps every request with a ULID-based id (prefix "req_"). When
//     trustUpstream is true the client-supplied header is accepted after a
//     cheap syntactic check (alnum+hyphen, <=128 chars); otherwise a fresh
//     id is generated.
//   - RequestIDFromContext(c) — helper used by other middlewares (413/429
//     envelopes, panic recovery) to echo the current request id.
//   - CtxKeyRequestID        — the gin + context.Context key shared by the
//     logger block so slog records inherit request_id automatically.
//
// The middleware must be the FIRST r.Use(...) so downstream middlewares can
// rely on the context key being populated.
const requestIDSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=request-id
//` + `ff:what RequestID — 모든 요청에 ULID 기반 request_id 부여 (X-Request-Id 응답 헤더 + context)

package middleware

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

// CtxKeyRequestID is the gin + context.Context key used to propagate the
// request id to logger handlers and inner middlewares. Exported as a typed
// string so collisions with user keys are impossible.
type ctxKeyRequestIDType string

// CtxKeyRequestID is the shared key. The string value matches the header
// name lowercased without dashes so slog auto-injection reads naturally.
const CtxKeyRequestID ctxKeyRequestIDType = "request_id"

// GinKeyRequestID is the gin.Context key (stored via c.Set). gin keys are
// plain strings; we keep both constants aligned for ergonomics.
const GinKeyRequestID = "request_id"

// RequestIDPrefix is prepended to all server-generated ids so logs show a
// single-glance marker distinct from DB / user ids.
const RequestIDPrefix = "req_"

// maxRequestIDLen bounds accepted upstream ids. 128 covers UUID (36),
// ULID (26), KSUID (27), and tracing frameworks that concatenate trace +
// span ids (e.g. W3C traceparent 55 bytes) with headroom.
const maxRequestIDLen = 128

// ulidEntropyPool supplies rand.Source-backed entropy readers to
// ulid.New without requiring a lock on every call. Each worker goroutine
// receives its own reader through a sync.Pool.
var ulidEntropyPool = sync.Pool{
	New: func() interface{} {
		return ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	},
}

// RequestID returns a gin middleware that ensures every request carries a
// request id. trustUpstream=true accepts a client-supplied header after a
// cheap validity check; false always issues a server-generated id.
//
// The id is written to:
//   - c.Set(GinKeyRequestID, id)
//   - context.Context via c.Request = c.Request.WithContext(...)
//   - response header (header arg, default "X-Request-Id")
func RequestID(trustUpstream bool, header string) gin.HandlerFunc {
	if header == "" {
		header = "X-Request-Id"
	}
	return func(c *gin.Context) {
		id := ""
		if trustUpstream {
			id = sanitizeUpstreamID(c.GetHeader(header))
		}
		if id == "" {
			id = generateRequestID()
		}
		c.Set(GinKeyRequestID, id)
		ctx := context.WithValue(c.Request.Context(), CtxKeyRequestID, id)
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set(header, id)
		c.Next()
	}
}

// RequestIDFromContext returns the id stamped by RequestID on this gin
// context. Returns "" when RequestID has not run (e.g. pre-router panic).
func RequestIDFromContext(c *gin.Context) string {
	if v, ok := c.Get(GinKeyRequestID); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	if c.Request != nil {
		if v := c.Request.Context().Value(CtxKeyRequestID); v != nil {
			if s, ok := v.(string); ok {
				return s
			}
		}
	}
	return ""
}

// RequestIDFromStdContext is the net/http analogue used by panic recovery
// and graceful-shutdown hooks that only hold a context.Context.
func RequestIDFromStdContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if v := ctx.Value(CtxKeyRequestID); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// generateRequestID returns a fresh ULID string with the "req_" prefix.
// Uses the entropy pool so contention stays minimal under load.
func generateRequestID() string {
	entropy := ulidEntropyPool.Get().(*ulid.MonotonicEntropy)
	defer ulidEntropyPool.Put(entropy)
	id, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
	if err != nil {
		// Extremely unlikely (entropy exhaustion on the monotonic source);
		// fall back to a non-monotonic ULID. Still unique globally.
		return RequestIDPrefix + ulid.Make().String()
	}
	return RequestIDPrefix + id.String()
}

// sanitizeUpstreamID validates a client-supplied request id. Rules:
//   - length <= maxRequestIDLen
//   - characters: ASCII alnum, hyphen, underscore, dot, colon
//
// Returns the cleaned id or "" when the value is unacceptable.
func sanitizeUpstreamID(raw string) string {
	if raw == "" || len(raw) > maxRequestIDLen {
		return ""
	}
	for i := 0; i < len(raw); i++ {
		ch := raw[i]
		ok := (ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '_' || ch == '.' || ch == ':'
		if !ok {
			return ""
		}
	}
	return raw
}
`
