//ff:func feature=validate-contract type=rule control=iteration dimension=1 topic=preserve-safety
//ff:what prv15PreservedMapAccess — preserved 파일에서 map 접근 결과를 가드 없이 즉시 selector 사용 ERROR

package contract

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// prv15PreservedMapAccess runs PRV-15 against every preserved path.
// Per-file walking is delegated to scanFileForMapAccess which looks
// for inline `m[k].Field` or `m[k].Method()` patterns — the only
// shape we can flag without type information while keeping precision
// acceptable.
func prv15PreservedMapAccess(paths []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, path := range paths {
		diags = append(diags, scanFileForMapAccess(path)...)
	}
	return diags
}
