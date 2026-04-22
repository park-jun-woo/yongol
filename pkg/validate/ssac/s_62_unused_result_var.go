//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-62 — 선언된 result 변수가 후속 시퀀스에서 한 번도 참조되지 않으면 ERROR

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s62UnusedResultVar validates S-62: a result variable declared in a sequence
// must be referenced at least once in the sequences that follow it (Inputs
// values, Fields values, or Target). Emits an ERROR when never referenced.
func s62UnusedResultVar(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		seqs := fn.Sequences
		for i, seq := range seqs {
			if seq.Result == nil || seq.Result.Var == "" {
				continue
			}
			varName := seq.Result.Var
			if varName == "_" {
				continue
			}
			if s62unusedInSubsequent(varName, seqs, i+1) {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[S-62] result 변수 %q가 후속 시퀀스에서 사용되지 않습니다", varName),
					Advice:  "불필요한 변수는 제거하거나, 필요하다면 @response 등에서 참조하세요",
				})
			}
		}
	}
	return diags
}
