//ff:func feature=gen-fastapi type=generator control=iteration dimension=1
//ff:what renderInlineStubs — 같은 feature 내 @call 대상 inline stub 함수 렌더링

package ssac

import (
	"fmt"
	"strings"
)

// renderInlineStubs produces inline async stub functions for same-feature
// @call/@eval targets that lack definitions. Each stub raises
// NotImplementedError so the developer knows where to add real logic.
func renderInlineStubs(feature string, stubNames []string) string {
	if len(stubNames) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("\n# --- same-feature @call stubs (implement with actual business logic) ---\n\n")
	for _, fn := range stubNames {
		b.WriteString(fmt.Sprintf("async def %s(*args, **kwargs):\n", fn))
		b.WriteString(fmt.Sprintf("    raise NotImplementedError(\"%s.%s not implemented\")\n\n\n",
			feature, fn))
	}
	return b.String()
}
