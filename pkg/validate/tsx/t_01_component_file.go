//ff:func feature=validate type=rule control=iteration dimension=1 topic=tsx
//ff:what T-1 — 로컬 컴포넌트 import 경로의 실제 파일 존재 여부 확인
package tsx

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// t01ComponentFile validates T-1: for each PageSpec.Imports entry (local
// component imports only — npm packages are filtered out in the parser),
// the referenced file must exist on disk. WARNING level because AI agents
// iterate on TSX frequently and a transient missing file during scaffolding
// should not block the entire validate run.
func t01ComponentFile(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.TSXPages) == 0 {
		return nil
	}
	frontendRoot := filepath.Join(fs.SpecsDir, "frontend")
	aliasRoot := resolveAliasRoot(frontendRoot)

	var diags []diagnostic.Diagnostic
	for _, page := range fs.TSXPages {
		for _, imp := range page.Imports {
			target := resolveImportPath(imp.Path, page.File, aliasRoot)
			if target == "" {
				continue
			}
			if !fileExistsWithTsxVariants(target) {
				diags = append(diags, diagnostic.Diagnostic{
					File:    page.File,
					Line:    imp.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelWarning,
					Message: "[T-1] imported component file not found: " + imp.Path,
					Advice:  "해당 경로에 컴포넌트 파일(.tsx/.ts/index.tsx 중 하나)을 생성하거나 import 경로를 수정하세요",
				})
			}
		}
	}
	return diags
}

// resolveAliasRoot maps the @/ path-alias to a concrete directory.
// Convention:
//   - <frontend>/src exists  → @/ = <frontend>/src
//   - otherwise               → @/ = <frontend>
//
// This matches the Vite/tsconfig default where pages live under src/ but
// keeps legacy layouts (components directly under frontend/) working.
func resolveAliasRoot(frontendRoot string) string {
	if st, err := os.Stat(filepath.Join(frontendRoot, "src")); err == nil && st.IsDir() {
		return filepath.Join(frontendRoot, "src")
	}
	return frontendRoot
}

// resolveImportPath converts a TSX import source into an absolute filesystem
// path (without extension). Returns "" when the import is not a local one —
// npm packages should never reach this function since the parser filters
// them, but defensive check remains.
func resolveImportPath(importSource, fromFile, aliasRoot string) string {
	if strings.HasPrefix(importSource, "@/") {
		return filepath.Join(aliasRoot, strings.TrimPrefix(importSource, "@/"))
	}
	if strings.HasPrefix(importSource, "./") || strings.HasPrefix(importSource, "../") {
		return filepath.Join(filepath.Dir(fromFile), importSource)
	}
	return ""
}

// fileExistsWithTsxVariants returns true if any common resolution of `base`
// exists: base.tsx, base.ts, base.jsx, base.js, base/index.tsx, base/index.ts.
// The order mirrors Node's moduleResolution with TypeScript+JSX.
func fileExistsWithTsxVariants(base string) bool {
	// Direct file match (with extension already present).
	if _, err := os.Stat(base); err == nil {
		return true
	}
	suffixes := []string{".tsx", ".ts", ".jsx", ".js"}
	for _, s := range suffixes {
		if _, err := os.Stat(base + s); err == nil {
			return true
		}
	}
	indexes := []string{"index.tsx", "index.ts", "index.jsx", "index.js"}
	for _, idx := range indexes {
		if _, err := os.Stat(filepath.Join(base, idx)); err == nil {
			return true
		}
	}
	return false
}
