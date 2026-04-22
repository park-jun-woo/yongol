//ff:func feature=validate-contract type=test control=sequence
//ff:what TestCompareExternalSymbolsAllPresent — 모든 심볼이 SSOT에 존재할 때 drift 없음 검증

package contract

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

func TestCompareExternalSymbolsAllPresent(t *testing.T) {
	actual := contract.ExternalSymbols{
		SqlcQueries: []string{"UserFindByID"},
		CallTargets: []string{"billing.CheckCredits"},
		DDLFields:   []string{"u.Email"},
	}
	queries := map[string]bool{"UserFindByID": true}
	calls := map[string]bool{"billing.CheckCredits": true}
	fields := map[string]bool{"email": true}
	ms := compareExternalSymbols(actual, queries, calls, fields)
	if len(ms.Queries)+len(ms.Calls)+len(ms.Fields) != 0 {
		t.Errorf("expected no drift, got %+v", ms)
	}
}
