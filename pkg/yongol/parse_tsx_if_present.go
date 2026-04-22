//ff:func feature=orchestrator type=loader control=sequence
//ff:what TSX 탐지 시 ParseDir 실행 — 성공 시 TSXPages, 실패 시 parse diagnostic 수집
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/tsx"
)

// parseTSXIfPresent parses every .tsx file under the frontend/ directory
// when KindTSX is detected. swc/node failures surface as a single parse
// diagnostic so the CLI can fail fast with a clear install hint. Individual
// file parse errors are also aggregated into one diagnostic per failure;
// successfully parsed pages are still kept so partial extraction works in
// AI iteration loops.
func parseTSXIfPresent(fs *Fullstack, has map[SSOTKind]DetectedSSOT) {
	d, ok := has[KindTSX]
	if !ok {
		return
	}
	pages, err := tsx.ParseDir(d.Path)
	fs.TSXPages = pages
	if err != nil {
		fs.ParseDiagnostics = append(fs.ParseDiagnostics, diagnostic.Diagnostic{
			File:    d.Path,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "tsx parse failed: " + err.Error(),
			Advice:  "install @swc/core in the frontend project (npm install --save-dev @swc/core) or set YONGOL_SWC_PROJECT_DIR",
		})
	}
}
