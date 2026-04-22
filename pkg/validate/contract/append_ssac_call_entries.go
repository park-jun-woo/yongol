//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what appendSSaCCallEntries — SSaC.callRef 항목을 camelCase + PascalCase 두 형태로 calls 맵에 추가

package contract

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// appendSSaCCallEntries inserts each ref into calls verbatim, then also
// inserts the PascalCase variant of the name-after-dot so both call
// spellings resolve (`billing.checkCredits` and `billing.CheckCredits`).
func appendSSaCCallEntries(refs rule.StringSet, calls map[string]bool) {
	for ref := range refs {
		calls[ref] = true
		idx := strings.Index(ref, ".")
		if idx <= 0 || idx == len(ref)-1 {
			continue
		}
		pkg := ref[:idx]
		name := ref[idx+1:]
		calls[pkg+"."+strings.ToUpper(name[:1])+name[1:]] = true
	}
}
