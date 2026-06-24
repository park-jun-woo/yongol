//ff:func feature=orchestrator type=loader control=sequence
//ff:what 한 도메인의 frontend/layouts/ 를 os.Stat 가드 후 적재 — 단일 사이트 parseLayoutIfPresent 와 동일 시맨틱
package yongol

import (
	"os"
	"path/filepath"
)

// loadDomainLayouts loads one domain's optional frontend/layouts/ directory into
// fs.DomainLayouts. It is os.Stat-guarded and stores only on a clean parse,
// exactly like the single-site parseLayoutIfPresent loader.
func loadDomainLayouts(fs *Fullstack, name, frontDir string) {
	layoutDir := filepath.Join(frontDir, "layouts")
	info, err := os.Stat(layoutDir)
	if err != nil || !info.IsDir() {
		return
	}
	layouts, diags := parseLayouts(layoutDir)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.DomainLayouts[name] = layouts
	}
}
