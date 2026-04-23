//ff:func feature=gen-gogin type=util control=sequence
//ff:what generateGoMod — arts/backend/go.mod 생성 (go mod init + go get @latest + go mod tidy)
package gogin

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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

	if err := removeStaleGoModFiles(backendDir); err != nil {
		return err
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
	deps := resolveGoModDeps(fs)
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
	return runGoModTidy(backendDir)
}
