//ff:func feature=gen-fastapi type=util control=selection
//ff:what renderOneOp — 단일 Op → Python 문장 렌더링 (OpKind 분기)

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderOneOp dispatches rendering for a single Op to the type-specific renderer.
func renderOneOp(b *strings.Builder, op ir.Op, indent, sessionRef string) {
	switch op.Kind {
	case ir.OpGet:
		renderGetOp(b, op.Get, indent, sessionRef)
	case ir.OpPost:
		renderPostOp(b, op.Post, indent, sessionRef)
	case ir.OpPut:
		renderPutOp(b, op.Put, indent, sessionRef)
	case ir.OpDelete:
		renderDeleteOp(b, op.Delete, indent, sessionRef)
	case ir.OpEmpty:
		renderEmptyOp(b, op.Empty, indent)
	case ir.OpExists:
		renderExistsOp(b, op.Exists, indent)
	case ir.OpAuth:
		renderAuthOp(b, op.Auth, indent)
	case ir.OpState:
		renderStateOp(b, op.State, indent)
	case ir.OpCall:
		renderCallOp(b, op.Call, indent)
	case ir.OpEval:
		renderEvalOp(b, op.Eval, indent)
	case ir.OpPublish:
		renderPublishOp(b, op.Publish, indent)
	case ir.OpVerifyPassword:
		renderVerifyPasswordOp(b, op.VerifyPW, indent, sessionRef)
	case ir.OpResponse:
		renderResponseOp(b, op.Response, indent)
	}
}
