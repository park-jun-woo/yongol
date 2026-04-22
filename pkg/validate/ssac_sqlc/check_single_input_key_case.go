//ff:func feature=validate type=util control=sequence topic=sqlc
//ff:what checkSingleInputKeyCase — 단일 input key 의 sqlc param 대소문자 불일치 판정

package ssac_sqlc

import (
	"fmt"

	"github.com/ettle/strcase"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// checkSingleInputKeyCase reports (diag, true) when key does not match any
// sqlc params in its exact form nor in strcase.ToSnake form, but matches via
// case-insensitive lookup — indicating a casing mismatch only.
func checkSingleInputKeyCase(fn ssac.ServiceFunc, seq ssac.Sequence, key string, params map[string]bool) (diagnostic.Diagnostic, bool) {
	if params[key] {
		return diagnostic.Diagnostic{}, false
	}
	// Go convention: PascalCase struct fields ↔ snake_case sqlc params.
	// Accept `Email` ↔ `email`, `BidAmount` ↔ `bid_amount`, etc.
	if params[strcase.ToSnake(key)] {
		return diagnostic.Diagnostic{}, false
	}
	matched := findCaseInsensitiveParam(key, params)
	if matched == "" {
		return diagnostic.Diagnostic{}, false
	}
	return diagnostic.Diagnostic{
		File:    fn.FileName,
		Line:    seq.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: fmt.Sprintf("[XQS-14] input key %q와 sqlc 파라미터 %q 대소문자 불일치", key, matched),
		Advice:  fmt.Sprintf("input key를 %q로 변경하세요", matched),
	}, true
}
