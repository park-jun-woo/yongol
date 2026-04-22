//ff:func feature=validate type=util control=sequence topic=func-check
//ff:what callFuncName — @call "pkg.Func" 에서 PascalCase 함수명 부분 추출

package ssac_func

import "strings"

// callFuncName returns the PascalCase function name portion of a @call model.
// "auth.VerifyPassword" -> "VerifyPassword"
// Returns "" if the model is not a qualified pkg.Func form.
func callFuncName(model string) string {
	idx := strings.IndexByte(model, '.')
	if idx <= 0 || idx >= len(model)-1 {
		return ""
	}
	return model[idx+1:]
}
