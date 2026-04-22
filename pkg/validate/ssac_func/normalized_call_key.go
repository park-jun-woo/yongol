//ff:func feature=validate type=util control=sequence topic=func-check
//ff:what normalizedCallKey — SSaC @call "pkg.Pascal" → Func 어노테이션 "pkg.camel"

package ssac_func

import "strings"

// normalizedCallKey converts "billing.DeductCredit" (SSaC @call style) to
// "billing.deductCredit" (@func annotation style) for Func.spec / findFuncSpec
// lookups. Case-exact match after this normalization.
func normalizedCallKey(model string) string {
	idx := strings.IndexByte(model, '.')
	if idx < 0 {
		return model
	}
	return model[:idx+1] + toCamelKey(model[idx+1:])
}
