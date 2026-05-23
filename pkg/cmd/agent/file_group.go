//ff:type feature=agent type=helper
//ff:what fileGroup — 파일별로 그룹화된 진단 목록

package agent

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// fileGroup groups diagnostics by file path.
type fileGroup struct {
	relFile string // specs-dir relative path
	layer   layer
	diags   []diagnostic.Diagnostic
}
