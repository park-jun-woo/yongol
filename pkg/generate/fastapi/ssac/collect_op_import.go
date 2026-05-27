//ff:func feature=gen-fastapi type=util control=selection dimension=1
//ff:what collectOpImport — 단일 Op → importData 에 모델/DML/외부 패키지 참조 추가

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// collectOpImport processes a single op into the importData.
func collectOpImport(d *importData, op ir.Op) {
	switch op.Kind {
	case ir.OpGet:
		d.UsesSelect = true
		if op.Get != nil {
			d.Models[pascalCase(op.Get.Model)] = true
		}
	case ir.OpPost:
		if op.Post != nil {
			d.Models[pascalCase(op.Post.Model)] = true
		}
	case ir.OpPut:
		d.UsesUpdate = true
		if op.Put != nil {
			d.Models[pascalCase(op.Put.Model)] = true
		}
	case ir.OpDelete:
		d.UsesDelete = true
		if op.Delete != nil {
			d.Models[pascalCase(op.Delete.Model)] = true
		}
	case ir.OpPublish:
		d.HasPublish = true
	case ir.OpAuth:
		d.HasAuth = true
	case ir.OpCall:
		if op.Call != nil && op.Call.Package != "" {
			addExtPkgRef(d, op.Call.Package, op.Call.Function)
		}
	case ir.OpEval:
		if op.Eval != nil && op.Eval.Package != "" {
			addExtPkgRef(d, op.Eval.Package, op.Eval.Function)
		}
	}
}
