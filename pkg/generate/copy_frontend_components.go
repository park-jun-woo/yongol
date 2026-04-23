//ff:func feature=generate type=util control=sequence
//ff:what copyFrontendComponents — specs/frontend/** (components, pages, 기타 .tsx/.ts/.css) 를 arts/frontend/src/ 아래로 복제
package generate

import (
	"fmt"
	"os"
	"path/filepath"
)

// copyFrontendComponents mirrors user-authored frontend sources from
// specs/frontend/** into arts/frontend/src/**.
//
// Scope:
//   - Any .tsx / .ts / .css file under specs/frontend/ is copied preserving
//     its relative path so `@/components/...` and `@/pages/...` aliases
//     resolve identically between specs and arts.
//   - yongol-managed subtrees (src/components/ui/ — emitted as primitives,
//     src/api.ts, src/types/, src/lib/) are NOT copied from specs; those
//     are fully owned by the generator.
//   - Missing specs/frontend/ directory is a silent no-op (projects without
//     TSX SSOT remain supported).
//
// Rationale — TSX is the SSOT, so users author files directly. The Vite
// build then reads src/, not specs/frontend/, so the copy bridges the two
// until an explicit "serve directly from specs/frontend/" mode lands.
func copyFrontendComponents(specsDir, artifactsDir string) error {
	if specsDir == "" {
		return nil
	}
	srcRoot := filepath.Join(specsDir, "frontend")
	info, err := os.Stat(srcRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("stat %s: %w", srcRoot, err)
	}
	if !info.IsDir() {
		return nil
	}
	dstRoot := filepath.Join(artifactsDir, "frontend", "src")
	return filepath.Walk(srcRoot, makeFrontendCopyWalker(srcRoot, dstRoot))
}
