//ff:func feature=rule type=loader control=sequence
//ff:what registerSSaCCallRef — SSaC call "pkg.FuncName" 을 "pkg.funcName" 으로 정규화해 등록
package ground

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// registerSSaCCallRef normalizes SSaC call "pkg.FuncName" (PascalCase) →
// "pkg.funcName" (camelCase) to match @func annotation form used in Func.spec,
// and inserts it into callRefs.
func registerSSaCCallRef(model string, callRefs rule.StringSet) {
	idx := strings.IndexByte(model, '.')
	if idx <= 0 || idx >= len(model)-1 {
		return
	}
	tail := model[idx+1:]
	if tail[0] >= 'A' && tail[0] <= 'Z' {
		tail = string(tail[0]+('a'-'A')) + tail[1:]
	}
	callRefs[model[:idx+1]+tail] = true
}
