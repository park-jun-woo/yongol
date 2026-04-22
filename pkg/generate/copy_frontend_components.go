//ff:func feature=generate type=util control=iteration dimension=1
//ff:what copyFrontendSources — specs/frontend/** (components, pages, 기타 .tsx/.ts/.css) 를 arts/frontend/src/ 아래로 복제
package generate

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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
	return filepath.Walk(srcRoot, func(path string, fi os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if fi.IsDir() {
			// Skip node_modules in case a stray copy landed.
			if fi.Name() == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		rel, err := filepath.Rel(srcRoot, path)
		if err != nil {
			return err
		}
		if isYongolManaged(rel) {
			return nil
		}
		if !isCopiedExtension(path) {
			return nil
		}
		dst := filepath.Join(dstRoot, rel)
		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			return fmt.Errorf("mkdir %s: %w", filepath.Dir(dst), err)
		}
		return copyUserComponentFile(path, dst)
	})
}

// isYongolManaged is true when the relative path (inside specs/frontend/)
// names a subtree that the generator owns and should never be overwritten
// from specs. Users are free to author there but any conflict is resolved
// in favor of the generator.
func isYongolManaged(rel string) bool {
	rel = filepath.ToSlash(rel)
	switch {
	case rel == "src/api.ts",
		strings.HasPrefix(rel, "src/types/"),
		strings.HasPrefix(rel, "src/lib/"),
		strings.HasPrefix(rel, "src/components/ui/"):
		return true
	}
	return false
}

// isCopiedExtension whitelists file types relevant to a React source tree.
// CSS is included because shadcn primitives rely on tailwind utility classes
// but users sometimes override with custom stylesheets.
func isCopiedExtension(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".tsx", ".ts", ".jsx", ".js", ".css", ".svg":
		return true
	}
	return false
}

func copyUserComponentFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open %s: %w", src, err)
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create %s: %w", dst, err)
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy %s → %s: %w", src, dst, err)
	}
	return nil
}
