//ff:func feature=gen-gogin type=generator control=sequence
//ff:what domainRequestValidatorSource — 도메인별 request_validator_<ident>.go 소스 렌더링

package middleware

import (
	"fmt"
	"strings"
)

// domainRequestValidatorSource renders the per-domain request_validator_<ident>.go
// source for domain mode (BUG-142). Each domain embeds its own
// openapi_<ident>.yaml and exposes RequestValidator<Title>() so the per-domain
// route group can mount validation against that domain's contract only. All
// package-level identifiers are domain-suffixed (openapiSpec<Title>,
// RequestValidator<Title>) and bypassPrefixes is a function-local slice — both
// avoid same-package redeclaration when multiple domains coexist in the single
// internal/middleware package (same collision class as Phase008 §3d). The
// validator body mirrors the single-site requestValidatorSource exactly
// (kin-openapi ValidateRequest, /health /ready /metrics bypass, neutral 400 /
// 413 envelopes) so behavior stays identical per domain.
func domainRequestValidatorSource(name string) string {
	ident := domainIdent(name)
	title := domainTitle(name)
	var b strings.Builder
	// Build the //ff: header pieces split to keep this file's own filefunc
	// annotation from being mistaken for the generated file's directives.
	fmt.Fprintf(&b, "//%s feature=runtime-middleware type=util control=sequence\n",
		"ff:func")
	fmt.Fprintf(&b, "//%s RequestValidator%s — %s 도메인 kin-openapi 런타임 요청 검증 미들웨어\n\n",
		"ff:what", title, ident)
	b.WriteString(`package middleware

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

`)
	fmt.Fprintf(&b, "//go:embed openapi_%s.yaml\n", ident)
	fmt.Fprintf(&b, "var openapiSpec%s []byte\n\n", title)
	fmt.Fprintf(&b, "// RequestValidator%s loads the embedded %s-domain OpenAPI spec once at\n", title, ident)
	b.WriteString(`// startup and returns a gin middleware that rejects requests whose payload
// violates that domain's schema constraints. It is mounted on the domain's
// route group (group.Use) rather than globally so each domain validates only
// against its own contract. Returns an error instead of panicking so main()
// can exit cleanly via os.Exit(1) with a structured log line.
`)
	fmt.Fprintf(&b, "func RequestValidator%s() (gin.HandlerFunc, error) {\n", title)
	b.WriteString(`	bypassPrefixes := []string{"/health", "/ready", "/metrics"}
	loader := openapi3.NewLoader()
`)
	fmt.Fprintf(&b, "\tdoc, err := loader.LoadFromData(openapiSpec%s)\n", title)
	b.WriteString(`	if err != nil {
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
`)
	return b.String()
}
