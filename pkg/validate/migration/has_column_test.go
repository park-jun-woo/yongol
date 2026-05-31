//ff:func feature=validate type=test control=iteration dimension=1 topic=migration-hints
//ff:what hasColumn — nil/empty/존재/비존재 컬럼 검색 검증

package migration

import (
	"testing"

	gmig "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestHasColumn(t *testing.T) {
	tbl := &gmig.Table{
		Name: "users",
		Columns: []*gmig.Column{
			{Name: "id"},
			{Name: "email"},
			{Name: "name"},
		},
	}

	tests := []struct {
		name  string
		table *gmig.Table
		col   string
		want  bool
	}{
		{
			name:  "existing column found",
			table: tbl,
			col:   "id",
			want:  true,
		},
		{
			name:  "existing column in middle",
			table: tbl,
			col:   "email",
			want:  true,
		},
		{
			name:  "nonexistent column",
			table: tbl,
			col:   "phone",
			want:  false,
		},
		{
			name:  "empty columns list",
			table: &gmig.Table{Name: "empty"},
			col:   "id",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasColumn(tt.table, tt.col)
			if got != tt.want {
				t.Errorf("hasColumn(%q) = %v, want %v", tt.col, got, tt.want)
			}
		})
	}
}
