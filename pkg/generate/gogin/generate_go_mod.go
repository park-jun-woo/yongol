//ff:func feature=gen-gogin type=util control=sequence
//ff:what generateGoMod — arts/backend/go.mod 생성 (go mod init + go get @latest + go mod tidy)
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
// message. `go mod tidy` and friends can emit multi-KB diagnostics;
// truncating keeps the error readable while still surfacing the first
// concrete cause.
const goModTidyStderrLimit = 4 * 1024

// OTel dependency versions pinned centrally so otel-core, otelgin, otelsql,
// and the exporters stay on a compatible set. Bumps go through a single
// sweep rather than scattered literals across multiple generator blocks.
// Only this co-release family is pinned — everything else uses @latest and
// is frozen per-project in go.sum after the first generate.
const (
	otelVersion       = "v1.32.0"
	otelContribGinVer = "v0.57.0"
	otelSQLVer        = "v0.37.0"
)

// coreDeps are the non-OTel runtime dependencies of the generated backend.
// They use @latest at `go get` time — semver-major pins live in the import
// path (e.g. `jwt/v5`), so @latest only picks minor/patch upgrades. The
// generated project's own go.sum freezes the exact resolution per build.
var coreDeps = []string{
	"github.com/gin-gonic/gin@latest",
	"github.com/gin-contrib/cors@latest",
	"github.com/lib/pq@latest",
	"github.com/golang-jwt/jwt/v5@latest",
	"github.com/oapi-codegen/runtime@latest",
	"github.com/getkin/kin-openapi@latest",
	"github.com/ulule/limiter/v3@latest",
	"github.com/prometheus/client_golang@latest",
	"github.com/oklog/ulid/v2@latest",
	"github.com/park-jun-woo/ssac@latest",
}

// generateGoMod bootstraps arts/backend/go.mod using the Go standard
// toolchain: `go mod init`, `go get` for each dependency, optional
// `go mod edit -replace` for local ssac development, and `go mod tidy`
// to finalise indirect/go.sum. Any step failure propagates with captured
// stderr so callers surface the root cause instead of producing a
// silently-broken backend artifact.
//
// The previous implementation hand-assembled a `require (...)` block from
// hardcoded version literals, which let unresolvable pseudo-versions leak
// into the output. Delegating to `go get @latest` removes yongol's role as
// a package manager — semver-major compatibility is enforced by import
// paths and cross-project reproducibility is frozen in each project's
// go.sum.
//
// When manifest.backend.observability.tracing.enabled is true, the OTel
// dependency set is added. Exporter-specific modules (otlptracegrpc /
// stdouttrace) are fetched only when the matching exporter is configured,
// keeping the dependency surface minimal for projects that pick a single
// exporter at codegen time.
func generateGoMod(fs *yongol.Fullstack, module, artifactsDir string) error {
	backendDir := filepath.Join(artifactsDir, "backend")
	if err := os.MkdirAll(backendDir, 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", backendDir, err)
	}

	// Remove any stale go.mod/go.sum from a previous run — `go mod init`
	// refuses to overwrite an existing module file, and a leftover go.sum
	// would keep entries from a previous dependency set alive.
	for _, name := range []string{"go.mod", "go.sum"} {
		p := filepath.Join(backendDir, name)
		if err := os.Remove(p); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("remove stale %s: %w", name, err)
		}
	}

	// 1) Skeleton: module declaration + pinned go directive.
	if err := runGo(backendDir, "mod", "init", module); err != nil {
		return err
	}
	if err := runGo(backendDir, "mod", "edit", "-go=1.25.0"); err != nil {
		return err
	}

	// 2) Register dependencies via `go get`. @latest resolves within the
	// current major line (major is part of the import path in Go modules),
	// so breaking-change semver bumps cannot reach the build without an
	// explicit import-path change.
	deps := append([]string(nil), coreDeps...)
	if tracing := tracingEnabled(fs); tracing != nil {
		deps = append(deps,
			"go.opentelemetry.io/otel@"+otelVersion,
			"go.opentelemetry.io/otel/sdk@"+otelVersion,
			"go.opentelemetry.io/otel/trace@"+otelVersion,
			"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin@"+otelContribGinVer,
			"github.com/XSAM/otelsql@"+otelSQLVer,
		)
		switch tracing.Exporter {
		case "otlp", "":
			deps = append(deps, "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc@"+otelVersion)
		case "stdout":
			deps = append(deps, "go.opentelemetry.io/otel/exporters/stdout/stdouttrace@"+otelVersion)
		}
	}
	getArgs := append([]string{"get"}, deps...)
	if err := runGo(backendDir, getArgs...); err != nil {
		return err
	}

	// 3) Local development override — if YONGOL_LOCAL_SSAC points at a
	// ssac checkout, rewrite the ssac import via `go mod edit -replace` so
	// yongol authors can iterate on ssac/pkg/auth without pushing every
	// commit upstream before dummy projects pick it up. Production builds
	// leave the env unset and resolve via the module proxy.
	if localSSAC := strings.TrimSpace(os.Getenv("YONGOL_LOCAL_SSAC")); localSSAC != "" {
		if err := runGo(backendDir, "mod", "edit",
			"-replace=github.com/park-jun-woo/ssac="+localSSAC); err != nil {
			return err
		}
	}

	// 4) Finalise indirect requirements and go.sum.
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

// runGo executes `go <args...>` in dir, capturing stderr for error surfacing.
// Any non-zero exit wraps with the (truncated) stderr so callers see
// the concrete failure (network error, unresolved module, malformed flag)
// rather than a bare exit status.
func runGo(dir string, args ...string) error {
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go %s failed: %w: %s",
			strings.Join(args, " "), err,
			truncateStderr(stderr.String(), goModTidyStderrLimit))
	}
	return nil
}

// runGoModTidy executes `go mod tidy` in dir and captures stderr. On failure
// it returns a wrapped error carrying the (truncated) stderr so callers can
// surface the concrete cause (network error, auth failure, unresolved
// module, etc.) instead of silently producing a broken go.mod.
func runGoModTidy(dir string) error {
	return runGo(dir, "mod", "tidy")
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
