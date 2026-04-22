//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what s36CheckResponseStale — @response 필드에서 stale 변수 참조 검출 → WARNING

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// s36CheckResponseStale emits WARNING diagnostics for @response sequences
// that reference variables currently marked stale (mutated but not re-queried).
func s36CheckResponseStale(fn parsessac.ServiceFunc, _ int, seq parsessac.Sequence, stale map[string]bool) []diagnostic.Diagnostic {
	if seq.Type != "response" || seq.SuppressWarn {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, varRef := range seq.Fields {
		ref := inputValueRefBase(varRef)
		if ref == "" || !stale[ref] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    fn.FileName,
			Line:    seq.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: fmt.Sprintf("[S-36] @response uses %s which was mutated but not re-queried", ref),
			Advice:  "@put/@delete 후 변경된 객체를 @get 으로 다시 조회하세요",
		})
	}
	return diags
}
