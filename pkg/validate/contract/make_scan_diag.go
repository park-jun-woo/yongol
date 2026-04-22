//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what makeScanDiag — PRV-13 Diagnostic 값 생성 (메시지 포매팅 공통화)

package contract

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// makeScanDiag centralises PRV-13 Diagnostic creation, symmetrically
// with makeUnmarshalDiag. Keeping the message text in one place
// means future Rule-ID / wording changes touch a single file.
func makeScanDiag(path string, line int) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    path,
		Line:    line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[PRV-13] preserved file ignores sql.Scan error (line %d)", line),
		Advice: "Check the returned error before using the scanned fields:\n" +
			"  if err := row.Scan(&a, &b); err != nil { return api.Error500, err }\n" +
			"Add `// nolint:prv-13` for an intentional exception.",
	}
}
