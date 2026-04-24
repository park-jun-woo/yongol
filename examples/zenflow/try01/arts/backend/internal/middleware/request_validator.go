//ff:func feature=runtime-middleware type=util control=sequence
//ff:what RequestValidator — kin-openapi 기반 런타임 요청 검증 미들웨어
//ff:checked llm=yongol-gen hash=6839d49c

package middleware

import (
	_ "embed"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gin-gonic/gin"
)

//go:embed openapi.yaml
var openapiSpec []byte

// bypassPrefixes lists path prefixes that should skip validation entirely.
// health/ready probes and metrics endpoints are registered directly on the
// gin engine and have no OpenAPI route — validator lookup would 404 on them
// and, even with early-return on lookup miss, skipping is clearer.
var bypassPrefixes = []string{"/health", "/ready", "/metrics"}

// RequestValidator loads the embedded OpenAPI spec once at startup and
// returns a gin middleware that rejects requests whose payload violates
// schema constraints. Validation failures produce HTTP 400 with a neutral
// "Bad request" error and kin-openapi's native message as details.
// Returns an error instead of panicking on bootstrap failure so main() can
// exit cleanly via os.Exit(1) with a structured log line.
func RequestValidator() (gin.HandlerFunc, error) {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(openapiSpec)
	if err != nil {
		return nil, fmt.Errorf("load openapi spec: %w", err)
	}
	if err := doc.Validate(loader.Context); err != nil {
		return nil, fmt.Errorf("validate openapi spec: %w", err)
	}
	router, err := gorillamux.NewRouter(doc)
	if err != nil {
		return nil, fmt.Errorf("build openapi router: %w", err)
	}
	return func(c *gin.Context) {
		for _, p := range bypassPrefixes {
			if strings.HasPrefix(c.Request.URL.Path, p) {
				c.Next()
				return
			}
		}
		route, params, err := router.FindRoute(c.Request)
		if err != nil {
			// No matching route — let gin's own 404 path handle it.
			c.Next()
			return
		}
		input := &openapi3filter.RequestValidationInput{
			Request:    c.Request,
			PathParams: params,
			Route:      route,
			Options: &openapi3filter.Options{
				AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
			},
		}
		if err := openapi3filter.ValidateRequest(c.Request.Context(), input); err != nil {
			var mbe *http.MaxBytesError
			if errors.As(err, &mbe) {
				slog.Warn("payload too large", "path", c.Request.URL.Path, "limit", mbe.Limit)
				WriteEnvelopeWithContext(c,
					http.StatusRequestEntityTooLarge,
					"payload_too_large",
					"",
					map[string]interface{}{"limit": mbe.Limit},
				)
				return
			}
			slog.Warn("request validation failed", "path", c.Request.URL.Path, "err", err)
			WriteEnvelope(c, http.StatusBadRequest, "bad_request", err.Error())
			return
		}
		c.Next()
	}, nil
}
