//ff:func feature=orchestrator type=loader control=sequence
//ff:what STML 존재 시 layouts/ 하위 디렉토리에서 레이아웃 파싱 — 진단은 수집, 성공 시 Layouts 설정
package yongol

import (
	"os"
	"path/filepath"
)

// parseLayoutIfPresent parses the layouts/ subdirectory under the STML
// frontend directory when it exists. Called after STML page parsing.
// d.Path for KindSTML points to the specs/frontend directory.
func parseLayoutIfPresent(fs *Fullstack, has map[SSOTKind]DetectedSSOT) {
	d, ok := has[KindSTML]
	if !ok {
		return
	}
	layoutDir := filepath.Join(d.Path, "layouts")

	info, err := os.Stat(layoutDir)
	if err != nil || !info.IsDir() {
		return
	}

	layouts, diags := parseLayouts(layoutDir)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.Layouts = layouts
	}
}
