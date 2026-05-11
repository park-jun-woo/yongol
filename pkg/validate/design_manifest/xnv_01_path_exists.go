//ff:func feature=validate type=rule control=sequence topic=design-manifest
//ff:what XNV-01 — manifest.frontend.design 경로가 실제 파일로 존재하는지 확인
package design_manifest

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xnv01PathExists enforces XNV-01 (ERROR): when manifest.frontend.design
// declares a path, that path must resolve to an existing file relative to
// the specs directory.
func xnv01PathExists(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	designPath := fs.Manifest.Frontend.Design
	if designPath == "" {
		return nil
	}

	abs := filepath.Join(fs.SpecsDir, designPath)
	if _, err := os.Stat(abs); err == nil {
		return nil
	}

	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[XNV-01] frontend.design path %q does not exist", designPath),
		Advice:  "Create the DESIGN.md file at the declared path or fix the path in manifest.yaml.",
	}}
}
