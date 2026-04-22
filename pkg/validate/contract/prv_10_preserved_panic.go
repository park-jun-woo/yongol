//ff:func feature=validate-contract type=rule control=iteration dimension=1 topic=preserve-safety
//ff:what prv10PreservedPanic — preserved 파일 함수 body (init 제외) 의 panic() 호출 ERROR

package contract

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// prv10PreservedPanic runs PRV-10 against every preserved path. Per-
// file scanning is delegated to scanFileForPanic so this orchestrator
// stays flat and within Q4 body-length limits. Allowlist handling
// (init() body, `// nolint:panic`) lives in the scanner.
func prv10PreservedPanic(paths []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, path := range paths {
		diags = append(diags, scanFileForPanic(path)...)
	}
	return diags
}
