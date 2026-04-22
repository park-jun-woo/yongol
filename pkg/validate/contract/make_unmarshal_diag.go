//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what makeUnmarshalDiag — PRV-12 Diagnostic 값 생성 (메시지 포매팅 공통화)

package contract

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// makeUnmarshalDiag centralises PRV-12 Diagnostic creation so the
// orchestrating scanner stays short. Keeping message text in one
// place also simplifies future i18n or Rule-ID refactors.
func makeUnmarshalDiag(path string, line int) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    path,
		Line:    line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[PRV-12] preserved file ignores Unmarshal error (line %d)", line),
		Advice: "Check the returned error before using the decoded value:\n" +
			"  if err := json.Unmarshal(body, &req); err != nil {\n" +
			"      return api.Error400, nil\n" +
			"  }\n" +
			"If intentional, discard explicitly: `_ = json.Unmarshal(...)` or add `// nolint:prv-12`.",
	}
}
