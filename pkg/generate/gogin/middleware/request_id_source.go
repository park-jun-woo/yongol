//ff:type feature=gen-gogin type=generator
//ff:what requestIDSources — request_id 를 6파일로 분할하는 소스 템플릿 맵

package middleware

// requestIDTypeSource — ctxKeyRequestIDType type + constants + pool.
const requestIDTypeSource = `//` + `ff:type feature=runtime-middleware type=model topic=request-id
//` + `ff:what ctxKeyRequestIDType — request_id context key 타입 + 관련 상수/풀

package middleware

import (
	"math/rand"
	"sync"
	"time"

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
`

// requestIDMainSource — RequestID func.
const requestIDMainSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=request-id
//` + `ff:what RequestID — 모든 요청에 ULID 기반 request_id 부여 (X-Request-Id 응답 헤더 + context)

package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

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
`

// requestIDFromContextSource — RequestIDFromContext func.
const requestIDFromContextSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=request-id
//` + `ff:what RequestIDFromContext — gin.Context 에서 request_id 추출

package middleware

import "github.com/gin-gonic/gin"

// RequestIDFromContext returns the id stamped by RequestID on this gin
// context. Returns "" when RequestID has not run (e.g. pre-router panic).
func RequestIDFromContext(c *gin.Context) string {
	if v, ok := c.Get(GinKeyRequestID); ok {
		s, _ := v.(string)
		if s != "" {
			return s
		}
	}
	if c.Request == nil {
		return ""
	}
	v := c.Request.Context().Value(CtxKeyRequestID)
	if v == nil {
		return ""
	}
	s, _ := v.(string)
	return s
}
`

// requestIDFromStdContextSource — RequestIDFromStdContext func.
const requestIDFromStdContextSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=request-id
//` + `ff:what RequestIDFromStdContext — context.Context 에서 request_id 추출 (net/http 전용)

package middleware

import "context"

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
`

// generateRequestIDSource — generateRequestID func.
const generateRequestIDSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=request-id
//` + `ff:what generateRequestID — ULID 기반 "req_" prefix request_id 생성

package middleware

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// generateRequestID returns a fresh ULID string with the "req_" prefix.
// Uses the entropy pool so contention stays minimal under load.
func generateRequestID() string {
	entropy := ulidEntropyPool.Get().(*ulid.MonotonicEntropy)
	defer ulidEntropyPool.Put(entropy)
	id, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
	if err != nil {
		return RequestIDPrefix + ulid.Make().String()
	}
	return RequestIDPrefix + id.String()
}
`

// sanitizeUpstreamIDSource — sanitizeUpstreamID func.
const sanitizeUpstreamIDSource = `//` + `ff:func feature=runtime-middleware type=util control=iteration dimension=1 topic=request-id
//` + `ff:what sanitizeUpstreamID — 클라이언트 request_id 유효성 검증 (길이 + 허용 문자)

package middleware

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
