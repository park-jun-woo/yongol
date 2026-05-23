//ff:func feature=validate type=test control=iteration dimension=1 topic=migration-hints
//ff:what Mig001RenameMismatch — nil hints + 테이블/컬럼 rename 검증

package migration

import (
	"testing"

	gmig "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMig001RenameMismatch(t *testing.T) {
	prev := &gmig.Schema{
		Tables: map[string]*gmig.Table{
			"old_users": {
				Name: "old_users",
				Columns: []*gmig.Column{
					{Name: "old_col"},
				},
			},
		},
	}
	curr := &gmig.Schema{
		Tables: map[string]*gmig.Table{
			"new_users": {
				Name: "new_users",
				Columns: []*gmig.Column{
					{Name: "new_col"},
				},
			},
		},
	}

	tests := []struct {
		name      string
		hints     *gmig.Hints
		wantCount int
	}{
		{
			name:      "nil hints returns nil",
			hints:     nil,
			wantCount: 0,
		},
		{
			name:      "empty hints returns nil",
			hints:     &gmig.Hints{},
			wantCount: 0,
		},
		{
			name: "valid table rename no diagnostics",
			hints: &gmig.Hints{
				RenameTables: []gmig.RenameTableHint{
					{From: "old_users", To: "new_users"},
				},
			},
			wantCount: 0,
		},
		{
			name: "invalid table rename raises diagnostics",
			hints: &gmig.Hints{
				RenameTables: []gmig.RenameTableHint{
					{From: "nonexistent", To: "also_nonexistent"},
				},
			},
			wantCount: 2,
		},
		{
			name: "mixed valid table and invalid column rename",
			hints: &gmig.Hints{
				RenameTables: []gmig.RenameTableHint{
					{From: "old_users", To: "new_users"},
				},
				RenameColumns: []gmig.RenameColumnHint{
					{Table: "old_users", From: "bad_col", To: "new_col"},
				},
			},
			wantCount: 1, // from=bad_col not in prev's old_users
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := Mig001RenameMismatch(prev, curr, tt.hints)
			if len(diags) != tt.wantCount {
				t.Fatalf("expected %d diagnostics, got %d: %+v", tt.wantCount, len(diags), diags)
			}
		})
	}
}
