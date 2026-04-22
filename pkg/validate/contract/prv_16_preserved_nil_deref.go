//ff:func feature=validate-contract type=rule control=iteration dimension=1 topic=preserve-safety
//ff:what prv16PreservedNilDeref — preserved 파일에서 Get*/Find* 호출 결과 즉시 selector 접근 ERROR

package contract

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// prv16PreservedNilDeref runs PRV-16 against every preserved path.
// Per-file walking is delegated to scanFileForNilDeref which targets
// the specific shape `GetX()` or `FindX()` followed by `.Field` — a
// common nil-deref trap in handler code where the lookup could
// plausibly return nil.
func prv16PreservedNilDeref(paths []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, path := range paths {
		diags = append(diags, scanFileForNilDeref(path)...)
	}
	return diags
}
