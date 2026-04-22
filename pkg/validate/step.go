//ff:type feature=validate type=model
//ff:what step — 단일 검증 단위 (SSOT 또는 pair 폴더) 메타데이터
package validate

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// step is a single validation unit — one SSOT folder or one pair folder.
// Run receives the fully wired Fullstack (including fs.Ground()); no separate
// ground parameter needed.
type step struct {
	Name  string
	Kinds []yongol.SSOTKind
	Run   func(fs *yongol.Fullstack) []diagnostic.Diagnostic
}
