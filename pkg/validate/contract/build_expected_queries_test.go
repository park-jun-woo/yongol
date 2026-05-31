//ff:func feature=validate-contract type=test control=sequence
//ff:what TestBuildExpectedQueries — SQLcQueries 의 Name/Method 모두 허용 집합에 포함 검증
package contract

import (
	"testing"
)

func TestBuildExpectedQueries(t *testing.T) {
	fs := buildFSForPRV02() // single query UserFindByID/FindByID
	q := buildExpectedQueries(fs)
	if !q["UserFindByID"] {
		t.Error("expected raw -- name: ident present")
	}
	if !q["FindByID"] {
		t.Error("expected prefix-stripped method present")
	}
	if len(q) != 2 {
		t.Errorf("expected 2 entries, got %d (%v)", len(q), q)
	}
}
