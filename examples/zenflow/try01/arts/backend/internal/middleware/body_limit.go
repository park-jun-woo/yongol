//ff:func feature=runtime-middleware type=util control=sequence topic=dos-guard
//ff:what BodyLimit / MultipartLimit — HTTP body 크기 제한 (DoS 방지, http.MaxBytesReader)
//ff:checked llm=yongol-gen hash=ce4d3dc4

package middleware

import (
	"errors"
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

func applyOverride(c *gin.Context, maxBytes int64) {
	if maxBytes <= 0 {
		return
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
}

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
