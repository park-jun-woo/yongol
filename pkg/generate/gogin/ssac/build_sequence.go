//ff:func feature=gen-gogin type=util control=selection
//ff:what buildSequence — SSaC 시퀀스 타입 분기 → methodGen 메서드 호출

package ssac

import (
	"fmt"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildSequence dispatches a single SSaC sequence to its type-specific builder.
// next is the immediate follow-up sequence (or nil if seq is last) so builders
// like buildGet can adapt error handling when @empty/@exists follows.
// Returns (lines, imports, isPostCommit, err). An error is returned when the
// sequence type is not recognised — this prevents silent fall-through that
// would produce a build-success artefact missing the requested behaviour.
func (g *methodGen) buildSequence(seq ssacparser.Sequence, next *ssacparser.Sequence) ([]string, []string, bool, error) {
	switch seq.Type {
	case "get":
		l, imp := g.buildGet(seq, next)
		return l, imp, false, nil
	case "post":
		l, imp := g.buildPost(seq)
		return l, imp, false, nil
	case "put":
		l, imp := g.buildPut(seq)
		return l, imp, false, nil
	case "delete":
		l, imp := g.buildDelete(seq)
		return l, imp, false, nil
	case "empty":
		l, imp := g.buildEmpty(seq)
		return l, imp, false, nil
	case "exists":
		l, imp := g.buildExists(seq)
		return l, imp, false, nil
	case "state":
		l, imp := g.buildState(seq)
		return l, imp, false, nil
	case "auth":
		l, imp := g.buildAuth(seq)
		return l, imp, false, nil
	case "call":
		l, imp := g.buildCall(seq)
		return l, imp, false, nil
	case "eval":
		l, imp := g.buildEval(seq)
		return l, imp, false, nil
	case "verify-password":
		l, imp := g.buildVerifyPassword(seq)
		return l, imp, false, nil
	case "publish":
		// Phase006: @publish is emitted inside the tx block so its enqueue
		// INSERT shares the business transaction (outbox atomicity). For
		// handlers without a tx (useTx=false) buildPublish still emits the
		// legacy queue.Publish call.
		l, imp := g.buildPublish(seq)
		return l, imp, false, nil
	case "response":
		l, imp := g.buildResponse(seq)
		return l, imp, false, nil
	default:
		return nil, nil, false, fmt.Errorf(
			"unhandled SSaC sequence type: %q (file=%s, line=%d, func=%s)",
			seq.Type, g.FileName, seq.Line, g.FuncName,
		)
	}
}
