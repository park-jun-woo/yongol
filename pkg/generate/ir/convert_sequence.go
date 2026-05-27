//ff:func feature=gen-ir type=util control=selection
//ff:what convertSequence -- SSaC Sequence 타입별 분기 → IR Op 변환 (fs 참조로 OpenAPI/DDL/Rego/StateDiagram 해석 정보 이식)

package ir

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// convertSequence converts a single SSaC Sequence into an IR Op. The
// Fullstack context is threaded through so converters that enrich the IR
// with OpenAPI/DDL/Rego/StateDiagram metadata can access the parsed SSOTs.
func convertSequence(seq ssac.Sequence, fs *yongol.Fullstack) (Op, error) {
	switch seq.Type {
	case ssac.SeqGet:
		return convertGet(seq, fs), nil
	case ssac.SeqPost:
		return convertPost(seq), nil
	case ssac.SeqPut:
		return convertPut(seq), nil
	case ssac.SeqDelete:
		return convertDelete(seq), nil
	case ssac.SeqEmpty:
		return convertEmpty(seq), nil
	case ssac.SeqExists:
		return convertExists(seq), nil
	case ssac.SeqAuth:
		return convertAuth(seq, fs), nil
	case ssac.SeqState:
		return convertState(seq, fs), nil
	case ssac.SeqCall:
		return convertCall(seq), nil
	case ssac.SeqEval:
		return convertEval(seq), nil
	case ssac.SeqPublish:
		return convertPublish(seq), nil
	case ssac.SeqVerifyPassword:
		return convertVerifyPassword(seq), nil
	case ssac.SeqResponse:
		return convertResponse(seq), nil
	default:
		return Op{}, fmt.Errorf("unknown sequence type: %q", seq.Type)
	}
}
