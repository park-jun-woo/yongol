//ff:func feature=validate-contract type=util control=sequence
//ff:what diagnoseMissingQuery — 사라진 sqlc 쿼리 참조에 대한 PRV-02 Diagnostic 생성

package contract

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// diagnoseMissingQuery returns a PRV-02 Diagnostic for a preserved file
// that still calls a sqlc query the SSOT no longer defines.
func diagnoseMissingQuery(path, query string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    path,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[PRV-02] preserved, but references sqlc query %q which is not in any specs/db/*.sql — SSOT removed it.", query),
		Advice: strings.Join([]string{
			"(a) restore the -- name: " + query + " entry in specs/db/",
			"(b) remove the call from the preserved function body",
			"(c) release preserve by deleting the file",
		}, "\n"),
	}
}
