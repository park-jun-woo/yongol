//ff:func feature=validate-contract type=rule control=iteration dimension=1 topic=preserve-safety
//ff:what prv11PreservedCurrentUserAssertion — preserved 파일의 currentUser 단일 대입 타입 단언 ERROR

package contract

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// prv11PreservedCurrentUserAssertion runs PRV-11 against every
// preserved path. Per-file scanning is delegated to
// scanFileForCurrentUserAssertion so the orchestrator stays flat.
func prv11PreservedCurrentUserAssertion(paths []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, path := range paths {
		diags = append(diags, scanFileForCurrentUserAssertion(path)...)
	}
	return diags
}
