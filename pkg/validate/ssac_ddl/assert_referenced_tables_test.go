//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=ssac-ddl
//ff:what assertReferencedTables — buildReferencedTables 결과 일치 검증 헬퍼

package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func assertReferencedTables(t *testing.T, funcs []ssac.ServiceFunc, want map[string]bool) {
	t.Helper()
	got := buildReferencedTables(funcs)
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d; got %v", len(got), len(want), got)
	}
	for k := range want {
		if !got[k] {
			t.Errorf("missing table %q", k)
		}
	}
}
