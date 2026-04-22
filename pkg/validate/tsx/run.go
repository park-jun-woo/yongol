//ff:func feature=validate type=rule control=sequence topic=tsx
//ff:what Run — TSX 자체 정합성 검증 (T-*) 실행 진입점
package tsx

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all TSX self-consistency rules. Currently only T-1; grows
// when additional per-page invariants (e.g. raw fetch detection) land.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, t01ComponentFile(fs)...)
	return diags
}
