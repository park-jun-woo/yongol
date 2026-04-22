//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what shorthandWrapperSkip — shorthand target이 Page/Cursor/[] 래퍼면 true

package openapi_ssac

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// shorthandWrapperSkip reports whether the shorthand @response target `varName`
// is bound to a Result whose Wrapper is non-empty (Page[T]/Cursor[T]/[]T).
// These cases are skipped per XSO-20 defeater.
func shorthandWrapperSkip(fn ssac.ServiceFunc, varName string) bool {
	for _, seq := range fn.Sequences {
		if seq.Result == nil || seq.Result.Var != varName {
			continue
		}
		if seq.Result.Wrapper != "" {
			return true
		}
	}
	return false
}
