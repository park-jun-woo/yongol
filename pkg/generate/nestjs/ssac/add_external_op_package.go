//ff:func feature=gen-nestjs type=util control=selection
//ff:what addExternalOpPackage — 단일 Op에서 외부 패키지명을 seen 맵에 추가

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// addExternalOpPackage adds the external package from a single op to the seen set.
func addExternalOpPackage(seen map[string]bool, op ir.Op) {
	switch op.Kind {
	case ir.OpCall:
		if op.Call != nil && op.Call.Package != "" {
			seen[op.Call.Package] = true
		}
	case ir.OpEval:
		if op.Eval != nil && op.Eval.Package != "" {
			seen[op.Eval.Package] = true
		}
	}
}
