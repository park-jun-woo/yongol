//ff:func feature=generate type=util control=sequence
//ff:what copyFrontendComponents — frontend SSOT 디렉토리(.tsx/.ts/.css)를 대상 src/ 아래로 복제 (소스·대상 디렉토리 명시 인자)
package generate

import (
	"fmt"
	"os"
	"path/filepath"
)

// copyFrontendComponents mirrors user-authored frontend sources from
// srcFrontendDir/** into dstSrcDir/**.
//
// Both directories are passed explicitly so the same helper serves the
// single-site call (src = <specs>/frontend, dst = <artifacts>/frontend/src)
// and the per-domain call (src = <specs>/<cfg.Frontend>, dst =
// <artifacts>/frontend/<domain>/src). The per-domain source is derived from
// fs.Manifest.Domains[name].Frontend by the caller because DomainView keeps
// SpecsDir shared — synthesizing the path here from SpecsDir+convention would
// bypass the view and not be domain-aware (Decision N).
//
// Scope:
//   - Any .tsx / .ts / .css file under srcFrontendDir is copied preserving its
//     relative path so `@/components/...` and `@/pages/...` aliases resolve
//     identically between specs and arts.
//   - yongol-managed subtrees (src/components/ui/ — emitted as primitives,
//     src/api.ts, src/types/, src/lib/) are NOT copied from specs; those are
//     fully owned by the generator.
//   - Missing/empty/non-dir srcFrontendDir is a silent no-op (projects without
//     TSX SSOT remain supported).
//
// Rationale — TSX is the SSOT, so users author files directly. The Vite build
// then reads src/, not the specs frontend dir, so the copy bridges the two
// until an explicit "serve directly from specs" mode lands.
func copyFrontendComponents(srcFrontendDir, dstSrcDir string) error {
	if srcFrontendDir == "" {
		return nil
	}
	info, err := os.Stat(srcFrontendDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("stat %s: %w", srcFrontendDir, err)
	}
	if !info.IsDir() {
		return nil
	}
	return filepath.Walk(srcFrontendDir, makeFrontendCopyWalker(srcFrontendDir, dstSrcDir))
}
