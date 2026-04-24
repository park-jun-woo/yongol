//ff:func feature=orchestrator type=loader control=sequence
//ff:what parseSsacInterfaces — ssac/pkg/*/interface.yaml 을 Fullstack.SsacInterfaces 로 로드

package yongol

import (
	"log/slog"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// parseSsacInterfaces loads every DB-using ssac package's interface.yaml into
// fs.SsacInterfaces, keyed by the package name declared inside the file. The
// map is the single source of truth used by both codegen (Phase002) and
// validate rules (Phase004/Phase005).
//
// Root resolution piggybacks on findYongolPkgRoot, which already returns the
// ssac/pkg directory under three fallbacks (env var, sibling clone, GOMODCACHE).
// When no root is found the map stays empty — callers treat absent interfaces
// as "package does not declare any DB ports" which is indistinguishable from a
// non-DB ssac package (e.g. crypto / slug / redact).
func parseSsacInterfaces(fs *Fullstack) {
	pkgRoot := findYongolPkgRoot()
	if pkgRoot == "" {
		slog.Debug("ssacmeta: ssac pkg root not found — skipping interface.yaml load")
		return
	}
	// ssacmeta.LoadPackageInterfaces expects the ssac repo root, while
	// findYongolPkgRoot returns its `pkg` child. Step up one level.
	ssacRoot := filepath.Dir(pkgRoot)
	ifaces, err := ssacmeta.LoadPackageInterfaces(ssacRoot)
	if err != nil {
		slog.Warn("ssacmeta: interface.yaml load failed", "root", ssacRoot, "err", err)
		return
	}
	fs.SsacInterfaces = ifaces
}
