//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestAppendSSaCCallEntries — SSaC.callRef 를 camelCase + PascalCase 두 형태로 추가 검증

package contract

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestAppendSSaCCallEntries(t *testing.T) {
	calls := map[string]bool{}
	refs := rule.StringSet{
		"billing.checkCredits": true,
		"nodot":                true, // no dot → only verbatim
		"trailing.":            true, // dot at end → only verbatim
	}
	appendSSaCCallEntries(refs, calls)

	if !calls["billing.checkCredits"] {
		t.Error("expected verbatim camelCase entry")
	}
	if !calls["billing.CheckCredits"] {
		t.Error("expected PascalCase variant entry")
	}
	if !calls["nodot"] {
		t.Error("expected verbatim no-dot entry")
	}
	if !calls["trailing."] {
		t.Error("expected verbatim trailing-dot entry")
	}
}
