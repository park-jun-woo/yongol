//ff:func feature=gen-gogin type=test-helper control=sequence
//ff:what assertFindDDLTableByModelName — findDDLTableByModelName 결과 테이블명 검증 헬퍼
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// assertFindDDLTableByModelName asserts findDDLTableByModelName returns the
// expected table (wantTable=="" means nil).
func assertFindDDLTableByModelName(t *testing.T, tables []ddl.Table, modelName, wantTable string) {
	t.Helper()
	got := findDDLTableByModelName(tables, modelName)
	if wantTable == "" {
		if got != nil {
			t.Errorf("expected nil, got %q", got.Name)
		}
		return
	}
	if got == nil || got.Name != wantTable {
		t.Errorf("findDDLTableByModelName(%q) = %v, want table %q", modelName, got, wantTable)
	}
}
