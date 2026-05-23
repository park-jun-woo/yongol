//ff:func feature=validate type=test-helper control=sequence
//ff:what assertFindUserTable — findUserTable 결과 nil/이름 assertion 헬퍼

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func assertFindUserTable(t *testing.T, fs *yongol.Fullstack, userTable string, wantNil bool, wantName string) {
	t.Helper()
	got := findUserTable(fs, userTable)
	if wantNil {
		if got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
		return
	}
	if got == nil {
		t.Fatal("expected non-nil result, got nil")
	}
	if got.Name != wantName {
		t.Errorf("got Name=%q, want %q", got.Name, wantName)
	}
}
