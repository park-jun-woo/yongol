//ff:func feature=gen-fastapi type=util control=selection
//ff:what addOpPackageRef — 단일 Op에서 외부 패키지 참조를 맵에 추가

package fastapi

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// addOpPackageRef adds an external package reference from a single op to the map.
func addOpPackageRef(pm map[string]map[string]bool, op ir.Op) {
	switch op.Kind {
	case ir.OpCall:
		if op.Call != nil && op.Call.Package != "" {
			ensurePkgMap(pm, op.Call.Package)
			pm[op.Call.Package][op.Call.Function] = true
		}
	case ir.OpEval:
		if op.Eval != nil && op.Eval.Package != "" {
			ensurePkgMap(pm, op.Eval.Package)
			pm[op.Eval.Package][op.Eval.Function] = true
		}
	}
}
