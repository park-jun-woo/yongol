//ff:type feature=gen-gogin type=generator
//ff:what bodyLimitSources — body_limit 를 5파일로 분할하는 소스 템플릿 맵

package middleware

// bodyLimitSource is the source for internal/middleware/body_limit.go.
const bodyLimitSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=dos-guard
//` + `ff:what BodyLimit — 비 multipart HTTP body 크기 제한 (DoS 방지, http.MaxBytesReader)

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// BodyLimit caps non-multipart request bodies to maxBytes. maxBytes <= 0
// disables the cap for this middleware instance.
//
// multipart/form-data is skipped so a dedicated MultipartLimit can gate
// upload endpoints at a higher ceiling. Overflow is caught after c.Next()
// via c.Errors (handlers / validators surface *http.MaxBytesError).
func BodyLimit(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if maxBytes > 0 && !strings.HasPrefix(c.ContentType(), "multipart/") {
			c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		}
		c.Next()
		respondIfBodyTooLarge(c)
	}
}
`

// multipartLimitSource is the source for internal/middleware/multipart_limit.go.
const multipartLimitSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=dos-guard
//` + `ff:what MultipartLimit — multipart/form-data body 크기 제한 (업로드 DoS 방지)

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// MultipartLimit caps multipart/form-data bodies (typically for file
// uploads). Non-multipart requests pass through.
func MultipartLimit(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if maxBytes > 0 && strings.HasPrefix(c.ContentType(), "multipart/") {
			c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		}
		c.Next()
		respondIfBodyTooLarge(c)
	}
}
`

// overrideBodyLimitSource is the source for internal/middleware/override_body_limit.go.
const overrideBodyLimitSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=dos-guard
//` + `ff:what OverrideBodyLimit — 라우트별 body limit 오버라이드 미들웨어

package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// OverrideBodyLimit tightens or lifts the body limit for specific routes.
// Keys are "METHOD PATH" where PATH matches c.FullPath() (gin route
// template). Both maps may be nil. A value <= 0 means "no limit" and
// restores the raw body (streaming endpoints opt out this way).
//
// Order matters: this middleware must run AFTER BodyLimit / MultipartLimit
// because the last MaxBytesReader installed on c.Request.Body wins.
func OverrideBodyLimit(bodyOverrides, multipartOverrides map[string]int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.Method + " " + c.FullPath()
		if strings.HasPrefix(c.ContentType(), "multipart/") {
			if v, ok := multipartOverrides[key]; ok {
				applyOverride(c, v)
			}
		} else {
			if v, ok := bodyOverrides[key]; ok {
				applyOverride(c, v)
			}
		}
		c.Next()
		respondIfBodyTooLarge(c)
	}
}
`

// applyOverrideSource is the source for internal/middleware/apply_override.go.
const applyOverrideSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=dos-guard
//` + `ff:what applyOverride — 단일 라우트의 body limit 오버라이드 적용

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// applyOverride replaces the request body with a MaxBytesReader capped at
// maxBytes. Called by OverrideBodyLimit for matched routes.
func applyOverride(c *gin.Context, maxBytes int64) {
	if maxBytes <= 0 {
		return
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
}
`

// respondIfBodyTooLargeSource is the source for internal/middleware/respond_if_body_too_large.go.
const respondIfBodyTooLargeSource = `//` + `ff:func feature=runtime-middleware type=util control=iteration dimension=1 topic=dos-guard
//` + `ff:what respondIfBodyTooLarge — c.Errors 에서 MaxBytesError 감지 시 413 응답

package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// respondIfBodyTooLarge inspects c.Errors for *http.MaxBytesError and, if
// the response has not yet been written, replaces it with a 413 payload
// envelope. This is the shared tail for all body-limit middlewares.
func respondIfBodyTooLarge(c *gin.Context) {
	if c.Writer.Written() {
		return
	}
	for _, ginErr := range c.Errors {
		var mbe *http.MaxBytesError
		if errors.As(ginErr.Err, &mbe) {
			WriteEnvelopeWithContext(c,
				http.StatusRequestEntityTooLarge,
				"payload_too_large",
				"",
				map[string]interface{}{"limit": mbe.Limit},
			)
			return
		}
	}
}
`
