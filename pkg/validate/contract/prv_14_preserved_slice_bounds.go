//ff:func feature=validate-contract type=rule control=iteration dimension=1 topic=preserve-safety
//ff:what prv14PreservedSliceBounds — preserved 파일에서 slice 첫 요소 접근 전 len 가드 누락 ERROR

package contract

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// prv14PreservedSliceBounds runs PRV-14 against every preserved path.
// Per-file scanning is delegated to scanFileForSliceBounds — it walks
// each function body, tracks observed `len(x)` guards, and flags
// direct `x[0]` access without a preceding guard.
func prv14PreservedSliceBounds(paths []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, path := range paths {
		diags = append(diags, scanFileForSliceBounds(path)...)
	}
	return diags
}
