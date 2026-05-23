//ff:func feature=validate type=test control=iteration dimension=1 topic=migration-hints
//ff:what mig001CheckRenameTable — from/to 테이블 존재/누락 진단 검증

package migration

import (
	"testing"

	gmig "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMig001CheckRenameTable(t *testing.T) {
	prev := &gmig.Schema{
		Tables: map[string]*gmig.Table{
			"old_users": {Name: "old_users"},
		},
	}
	curr := &gmig.Schema{
		Tables: map[string]*gmig.Table{
			"new_users": {Name: "new_users"},
		},
	}

	tests := []struct {
		name      string
		hint      gmig.RenameTableHint
		wantCount int
		wantSub   string
	}{
		{
			name:      "valid rename: from in prev, to in curr",
			hint:      gmig.RenameTableHint{From: "old_users", To: "new_users"},
			wantCount: 0,
		},
		{
			name:      "from missing in prev raises diagnostic",
			hint:      gmig.RenameTableHint{From: "nonexistent", To: "new_users"},
			wantCount: 1,
			wantSub:   "from=nonexistent",
		},
		{
			name:      "to missing in curr raises diagnostic",
			hint:      gmig.RenameTableHint{From: "old_users", To: "nonexistent"},
			wantCount: 1,
			wantSub:   "to=nonexistent",
		},
		{
			name:      "both from and to missing raises two diagnostics",
			hint:      gmig.RenameTableHint{From: "bad_from", To: "bad_to"},
			wantCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := mig001CheckRenameTable(prev, curr, tt.hint)
			assertDiagCount(t, diags, tt.wantCount, tt.wantSub)
		})
	}
}
