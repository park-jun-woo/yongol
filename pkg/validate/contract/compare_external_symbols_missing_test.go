//ff:func feature=validate-contract type=test control=sequence
//ff:what TestCompareExternalSymbolsMissingAll — 모든 카테고리 drift 감지 검증

package contract

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

func TestCompareExternalSymbolsMissingAll(t *testing.T) {
	actual := contract.ExternalSymbols{
		SqlcQueries: []string{"UserFindByID"},
		CallTargets: []string{"billing.CheckCredits"},
		DDLFields:   []string{"u.Email"},
	}
	// expCalls declares `billing.Other` so package `billing` is known to
	// the SSOT — `billing.CheckCredits` then surfaces as drift (missing
	// method), while an untracked-package call would be ignored entirely.
	expCalls := map[string]bool{"billing.Other": true}
	ms := compareExternalSymbols(actual, map[string]bool{}, expCalls, map[string]bool{})
	if len(ms.Queries) != 1 || ms.Queries[0] != "UserFindByID" {
		t.Errorf("queries drift not detected: %v", ms.Queries)
	}
	if len(ms.Calls) != 1 || ms.Calls[0] != "billing.CheckCredits" {
		t.Errorf("calls drift not detected: %v", ms.Calls)
	}
	if len(ms.Fields) != 1 || ms.Fields[0] != "u.Email" {
		t.Errorf("fields drift not detected: %v", ms.Fields)
	}
}
