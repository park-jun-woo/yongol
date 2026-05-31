//ff:func feature=gen-fastapi type=util control=selection
//ff:what collectFieldArgs — Op 에서 모든 FieldArg 슬라이스를 추출 (kind 불문)

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// collectFieldArgs returns every FieldArg contained in the given Op,
// regardless of its Kind. Op types without FieldArgs return nil.
func collectFieldArgs(op ir.Op) []ir.FieldArg {
	switch op.Kind {
	case ir.OpGet:
		if op.Get == nil {
			return nil
		}
		out := make([]ir.FieldArg, 0, len(op.Get.Args)+len(op.Get.PaginationArgs))
		out = append(out, op.Get.Args...)
		out = append(out, op.Get.PaginationArgs...)
		return out
	case ir.OpPost:
		if op.Post == nil {
			return nil
		}
		return op.Post.Args
	case ir.OpPut:
		if op.Put == nil {
			return nil
		}
		return op.Put.Args
	case ir.OpDelete:
		if op.Delete == nil {
			return nil
		}
		return op.Delete.Args
	case ir.OpAuth:
		if op.Auth == nil {
			return nil
		}
		return op.Auth.Inputs
	case ir.OpState:
		if op.State == nil {
			return nil
		}
		return op.State.Inputs
	case ir.OpCall:
		if op.Call == nil {
			return nil
		}
		return op.Call.Args
	case ir.OpEval:
		if op.Eval == nil {
			return nil
		}
		return op.Eval.Args
	case ir.OpPublish:
		if op.Publish == nil {
			return nil
		}
		out := make([]ir.FieldArg, 0, len(op.Publish.Payload)+len(op.Publish.Options))
		out = append(out, op.Publish.Payload...)
		out = append(out, op.Publish.Options...)
		return out
	default:
		return nil
	}
}
