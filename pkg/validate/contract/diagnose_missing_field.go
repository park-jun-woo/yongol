//ff:func feature=validate-contract type=util control=sequence
//ff:what diagnoseMissingField — 사라진 DDL 컬럼 필드 참조에 대한 PRV-02 Diagnostic 생성

package contract

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// diagnoseMissingField returns a PRV-02 Diagnostic for a preserved file
// that still accesses a struct field whose underlying DDL column was
// removed from the SSOT.
func diagnoseMissingField(path, selector string) diagnostic.Diagnostic {
	field := selector
	if idx := strings.LastIndex(selector, "."); idx >= 0 && idx < len(selector)-1 {
		field = selector[idx+1:]
	}
	return diagnostic.Diagnostic{
		File:    path,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[PRV-02] preserved, but references DDL-derived field %q (%s) which is not in any specs/db/ table — SSOT removed it.", field, selector),
		Advice: strings.Join([]string{
			"(a) restore the column in specs/db/*.sql",
			"(b) remove the field access from the preserved function body",
			"(c) release preserve by deleting the file",
		}, "\n"),
	}
}
