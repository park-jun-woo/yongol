//ff:func feature=validate-contract type=test control=sequence
//ff:what TestCompareExternalSymbols — 알려진/미지 패키지·denylist·일치 항목 필터링 검증
package contract

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

func TestCompareExternalSymbols(t *testing.T) {
	t.Run("unknown package call ignored", func(t *testing.T) {
		actual := contract.ExternalSymbols{
			CallTargets: []string{"unknownpkg.Foo"},
		}
		ms := compareExternalSymbols(actual, map[string]bool{}, map[string]bool{"billing.X": true}, map[string]bool{})
		if len(ms.Calls) != 0 {
			t.Errorf("unknown-package call should be ignored, got %v", ms.Calls)
		}
	})

	t.Run("denylisted query skipped", func(t *testing.T) {
		actual := contract.ExternalSymbols{
			SqlcQueries: []string{"WithTx", "Gone"},
		}
		ms := compareExternalSymbols(actual, map[string]bool{}, map[string]bool{}, map[string]bool{})
		if len(ms.Queries) != 1 || ms.Queries[0] != "Gone" {
			t.Errorf("expected only Gone flagged, got %v", ms.Queries)
		}
	})

	t.Run("all present → empty", func(t *testing.T) {
		actual := contract.ExternalSymbols{
			SqlcQueries: []string{"FindByID"},
			DDLFields:   []string{"u.Email"},
		}
		ms := compareExternalSymbols(actual,
			map[string]bool{"FindByID": true},
			map[string]bool{},
			map[string]bool{canonicalFieldKey("email"): true})
		if len(ms.Queries)+len(ms.Calls)+len(ms.Fields) != 0 {
			t.Errorf("expected no drift, got %+v", ms)
		}
	})
}
