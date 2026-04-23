//ff:func feature=validate type=util control=iteration dimension=1 topic=tsx
//ff:what T-1 헬퍼 — 단일 PageSpec.Imports 의 파일 존재 여부 검사

package tsx

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	tsxparser "github.com/park-jun-woo/yongol/pkg/parser/tsx"
)

// t01CheckPage scans a single PageSpec's imports and emits a WARNING for
// each local component import path that cannot be resolved on disk.
func t01CheckPage(page tsxparser.PageSpec, aliasRoot string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, imp := range page.Imports {
		target := resolveImportPath(imp.Path, page.File, aliasRoot)
		if target == "" {
			continue
		}
		if fileExistsWithTsxVariants(target) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    page.File,
			Line:    imp.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[T-1] imported component file not found: " + imp.Path,
			Advice:  "Create a component file (.tsx, .ts, or index.tsx) at that path, or correct the import path",
		})
	}
	return diags
}
