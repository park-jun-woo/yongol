//ff:func feature=orchestrator type=loader control=sequence
//ff:what DESIGN.md 탐지 시 ParseFile 실행 — 진단은 수집, 성공 시 DesignSpec 설정
package yongol

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

// parseDesignIfPresent parses DESIGN.md when a design path is declared in
// manifest.frontend.design or when frontend/DESIGN.md exists by convention.
func parseDesignIfPresent(fs *Fullstack, root string) {
	path := resolveDesignPath(fs, root)
	if path == "" {
		return
	}

	if _, err := os.Stat(path); err != nil {
		// File not found — Design SSOT is absent.
		return
	}

	fs.Presences[KindDesign] = SSOTPopulated

	spec, diags := design.ParseFile(path)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.DesignSpec = spec
	}
}

// resolveDesignPath determines the DESIGN.md path. Priority:
// 1. manifest.frontend.design (explicit path relative to specs root)
// 2. convention: frontend/DESIGN.md
func resolveDesignPath(fs *Fullstack, root string) string {
	if fs.Manifest != nil && fs.Manifest.Frontend.Design != "" {
		return filepath.Join(root, fs.Manifest.Frontend.Design)
	}
	// Convention fallback.
	return filepath.Join(root, "frontend", "DESIGN.md")
}
