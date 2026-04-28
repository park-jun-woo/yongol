//ff:func feature=gen-gogin type=generator control=iteration dimension=2
//ff:what EmitAll — fs.SsacInterfaces 순회하여 각 DB-using 패키지의 postgres.go 작성

package infra

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// EmitAll iterates over fs.SsacInterfaces and writes
// `<artifactsDir>/backend/internal/infra/<pkg>/postgres.go` for every
// DB-using ssac package whose declared ports include at least one active
// port under the current manifest.
//
// Inactive ports (`when:` evaluated false) are filtered; when every port of
// a package is inactive the package emits nothing (no file). This keeps the
// artifacts tree free of unused adapters when e.g. `manifest.cache.backend`
// is "memory".
//
// modulePath is the user project's Go module path and is embedded in the
// generated file's import for `<modulePath>/internal/db`.
func EmitAll(fs *yongol.Fullstack, artifactsDir, modulePath string) error {
	if fs == nil || len(fs.SsacInterfaces) == 0 {
		return nil
	}
	mctx := flattenManifest(fs)

	// Deterministic order.
	pkgs := make([]string, 0, len(fs.SsacInterfaces))
	for k := range fs.SsacInterfaces {
		pkgs = append(pkgs, k)
	}
	sort.Strings(pkgs)

	for _, pkg := range pkgs {
		if err := emitOnePackage(fs, pkg, mctx, modulePath, artifactsDir); err != nil {
			return fmt.Errorf("infra: emit %s: %w", pkg, err)
		}
	}
	return nil
}
