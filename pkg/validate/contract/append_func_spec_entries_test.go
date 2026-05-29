//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestAppendFuncSpecEntries — FuncSpec 목록을 "pkg.Name" 키로 calls 맵에 추가 검증

package contract

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestAppendFuncSpecEntries(t *testing.T) {
	calls := map[string]bool{}
	specs := []funcspec.FuncSpec{
		{Package: "auth", Name: "HashPassword"},
		{Package: "", Name: "Skipped"},      // empty package → skipped
		{Package: "billing", Name: ""},      // empty name → skipped
		{Package: "auth", Name: "VerifyPW"}, // second valid entry
	}
	appendFuncSpecEntries(specs, calls)
	if !calls["auth.HashPassword"] {
		t.Error("expected auth.HashPassword present")
	}
	if !calls["auth.VerifyPW"] {
		t.Error("expected auth.VerifyPW present")
	}
	if len(calls) != 2 {
		t.Errorf("expected exactly 2 entries, got %d (%v)", len(calls), calls)
	}
}
