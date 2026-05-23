//ff:func feature=validate type=test control=iteration dimension=1 topic=migration-hints
//ff:what mig001CheckRenameColumn — from/to 컬럼 존재/누락 진단 검증

package migration

import (
	"testing"

	gmig "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMig001CheckRenameColumn(t *testing.T) {
	prev := &gmig.Schema{
		Tables: map[string]*gmig.Table{
			"users": {
				Name: "users",
				Columns: []*gmig.Column{
					{Name: "old_name"},
					{Name: "email"},
				},
			},
		},
	}
	curr := &gmig.Schema{
		Tables: map[string]*gmig.Table{
			"users": {
				Name: "users",
				Columns: []*gmig.Column{
					{Name: "new_name"},
					{Name: "email"},
				},
			},
		},
	}

	tests := []struct {
		name      string
		hint      gmig.RenameColumnHint
		wantCount int
		wantSub   string
	}{
		{
			name:      "valid rename: from exists in prev, to exists in curr",
			hint:      gmig.RenameColumnHint{Table: "users", From: "old_name", To: "new_name"},
			wantCount: 0,
		},
		{
			name:      "from column missing in prev raises one diagnostic",
			hint:      gmig.RenameColumnHint{Table: "users", From: "nonexistent", To: "new_name"},
			wantCount: 1,
			wantSub:   "from=nonexistent",
		},
		{
			name:      "to column missing in curr raises one diagnostic",
			hint:      gmig.RenameColumnHint{Table: "users", From: "old_name", To: "nonexistent"},
			wantCount: 1,
			wantSub:   "to=nonexistent",
		},
		{
			name:      "both from and to missing raises two diagnostics",
			hint:      gmig.RenameColumnHint{Table: "users", From: "bad_from", To: "bad_to"},
			wantCount: 2,
		},
		{
			name:      "table not in prev: no from diagnostic",
			hint:      gmig.RenameColumnHint{Table: "orders", From: "old_name", To: "nonexistent"},
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := mig001CheckRenameColumn(prev, curr, tt.hint)
			assertDiagCount(t, diags, tt.wantCount, tt.wantSub)
		})
	}
}
