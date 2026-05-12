//ff:func feature=validate type=rule control=iteration dimension=1 topic=design-manifest
//ff:what XNV-02 — specs/frontend/ 내 DESIGN.md 파일이 manifest에 미선언이면 경고
package design_manifest

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xnv02Undeclared enforces XNV-02 (WARNING): when a DESIGN.md (or
// *.design.md) file exists under specs/frontend/ but is not declared
// in manifest.frontend.design, emit a warning.
func xnv02Undeclared(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}

	frontendDir := filepath.Join(fs.SpecsDir, "frontend")
	found := findDesignFiles(frontendDir)
	if len(found) == 0 {
		return nil
	}

	declared := fs.Manifest.Frontend.Design
	var diags []diagnostic.Diagnostic
	for _, f := range found {
		rel, _ := filepath.Rel(fs.SpecsDir, f)
		if rel == "" {
			rel = f
		}
		if declared != "" && normPath(declared) == normPath(rel) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    rel,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: fmt.Sprintf("[XNV-02] design file %q exists but is not declared in manifest.frontend.design", rel),
			Advice:  fmt.Sprintf("Add `design: %s` under frontend in manifest.yaml, or remove the file if unused.", rel),
		})
	}
	return diags
}
