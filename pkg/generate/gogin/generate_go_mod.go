//ff:func feature=gen-gogin type=util control=sequence
//ff:what generateGoMod — arts/backend/go.mod 생성 + go mod tidy 실행 (실패 시 에러 전파)
package gogin

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// goModTidyStderrLimit caps captured stderr bytes included in the error
// message. `go mod tidy` can emit multi-KB diagnostics; truncating keeps
// the error readable while still surfacing the first concrete cause.
const goModTidyStderrLimit = 4 * 1024

// OTel dependency versions pinned centrally so otel-core, otelgin, otelsql,
// and the exporters stay on a compatible set. Bumps go through a single
// sweep rather than scattered literals across multiple generator blocks.
const (
	otelVersion          = "v1.32.0"
	otelContribGinVer    = "v0.57.0"
	otelSQLVer           = "v0.37.0"
	otelSemconvMinor     = "v1.26.0"
)

// generateGoMod writes arts/backend/go.mod with the module path from manifest
// and a fixed set of require directives, then runs `go mod tidy`.
// A `go mod tidy` failure is propagated as an error (with captured stderr)
// so callers — and ultimately `yongol generate` — surface the root cause
// instead of producing a silently-broken backend artifact.
//
// When manifest.backend.observability.tracing.enabled is true, the OTel
// dependency set is appended to the require block. Exporter-specific
// modules (otlptracegrpc / stdouttrace) are added only when the matching
// exporter is configured, keeping the dependency surface minimal for
// projects that pick a single exporter at codegen time.
func generateGoMod(fs *yongol.Fullstack, module, artifactsDir string) error {
	backendDir := filepath.Join(artifactsDir, "backend")
	if err := os.MkdirAll(backendDir, 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", backendDir, err)
	}

	var b strings.Builder
	b.WriteString("module " + module + "\n\n")
	b.WriteString("go 1.25.0\n\n")
	b.WriteString("require (\n")
	b.WriteString("\tgithub.com/gin-gonic/gin v1.10.0\n")
	b.WriteString("\tgithub.com/gin-contrib/cors v1.7.2\n")
	b.WriteString("\tgithub.com/lib/pq v1.10.9\n")
	b.WriteString("\tgithub.com/park-jun-woo/ssac v0.0.0-20260422052142-cf846596b28e\n")
	b.WriteString("\tgithub.com/golang-jwt/jwt/v5 v5.2.1\n")
	b.WriteString("\tgithub.com/oapi-codegen/runtime v1.4.0\n")
	b.WriteString("\tgithub.com/getkin/kin-openapi v0.133.0\n")
	b.WriteString("\tgithub.com/ulule/limiter/v3 v3.11.2\n")
	b.WriteString("\tgithub.com/prometheus/client_golang v1.19.1\n")
	b.WriteString("\tgithub.com/oklog/ulid/v2 v2.1.0\n")

	if tracing := tracingEnabled(fs); tracing != nil {
		fmt.Fprintf(&b, "\tgo.opentelemetry.io/otel %s\n", otelVersion)
		fmt.Fprintf(&b, "\tgo.opentelemetry.io/otel/sdk %s\n", otelVersion)
		fmt.Fprintf(&b, "\tgo.opentelemetry.io/otel/trace %s\n", otelVersion)
		fmt.Fprintf(&b, "\tgo.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin %s\n", otelContribGinVer)
		fmt.Fprintf(&b, "\tgithub.com/XSAM/otelsql %s\n", otelSQLVer)
		switch tracing.Exporter {
		case "otlp", "":
			fmt.Fprintf(&b, "\tgo.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc %s\n", otelVersion)
		case "stdout":
			fmt.Fprintf(&b, "\tgo.opentelemetry.io/otel/exporters/stdout/stdouttrace %s\n", otelVersion)
		}
		_ = otelSemconvMinor // semconv is pinned via its own pseudo-version — go mod tidy selects it
	}

	b.WriteString(")\n")

	// Local development replace — if the env var YONGOL_LOCAL_SSAC points
	// at a checkout of github.com/park-jun-woo/ssac, the generated go.mod
	// rewrites the ssac import to that path. This lets yongol authors iterate
	// on ssac/pkg/auth without pushing every commit upstream before the
	// dummy projects can pick it up. Production builds leave the env unset
	// and resolve via the module proxy.
	if localSSAC := strings.TrimSpace(os.Getenv("YONGOL_LOCAL_SSAC")); localSSAC != "" {
		fmt.Fprintf(&b, "\nreplace github.com/park-jun-woo/ssac => %s\n", localSSAC)
	}

	modPath := filepath.Join(backendDir, "go.mod")
	if err := os.WriteFile(modPath, []byte(b.String()), 0o644); err != nil {
		return fmt.Errorf("write go.mod: %w", err)
	}

	if err := runGoModTidy(backendDir); err != nil {
		return err
	}
	return nil
}

// tracingEnabled returns the tracing block when enabled, else nil. Used by
// generateGoMod to gate OTel dependency injection and by downstream
// generators that need the exporter kind for require-list shaping.
func tracingEnabled(fs *yongol.Fullstack) *pmanifest.ObservabilityTracing {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	obs := fs.Manifest.Backend.Observability
	if obs == nil || obs.Tracing == nil {
		return nil
	}
	if !obs.Tracing.Enabled {
		return nil
	}
	return obs.Tracing
}

// runGoModTidy executes `go mod tidy` in dir and captures stderr. On failure
// it returns a wrapped error carrying the (truncated) stderr so callers can
// surface the concrete cause (network error, auth failure, unresolved
// module, etc.) instead of silently producing a broken go.mod.
func runGoModTidy(dir string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir
	// GOTOOLCHAIN not forced — uses default (auto or local 1.25+)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go mod tidy failed: %w: %s", err, truncateStderr(stderr.String(), goModTidyStderrLimit))
	}
	return nil
}

// truncateStderr trims stderr output to at most limit bytes, appending a
// marker when truncation occurs. Trailing whitespace is also stripped so
// error messages stay on a single line where possible.
func truncateStderr(s string, limit int) string {
	s = strings.TrimRight(s, "\n\r\t ")
	if limit <= 0 || len(s) <= limit {
		return s
	}
	return s[:limit] + "...(truncated)"
}
