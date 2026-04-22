//ff:func feature=validate-contract type=util control=sequence
//ff:what diagnoseMissingCall — 사라진 @call / func 대상 참조에 대한 PRV-02 Diagnostic 생성

package contract

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// diagnoseMissingCall returns a PRV-02 Diagnostic for a preserved file
// that invokes a `<pkg>.<Func>` target no longer present in the Func
// SSOT or in SSaC @call refs.
func diagnoseMissingCall(path, target string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    path,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[PRV-02] preserved, but references call target %q which has no Func spec or SSaC @call — SSOT removed it.", target),
		Advice: strings.Join([]string{
			"(a) restore the function in specs/func/ or reintroduce the @call in the SSaC file",
			"(b) remove the call from the preserved function body",
			"(c) release preserve by deleting the file",
		}, "\n"),
	}
}
