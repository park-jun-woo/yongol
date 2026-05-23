//ff:func feature=validate type=test control=iteration dimension=1 topic=migration-hints
//ff:what Mig003DataMigrationMissing — nil/empty/복수 누락 경로 진단 검증

package migration

import (
	"testing"
)

func TestMig003DataMigrationMissing(t *testing.T) {
	tests := []struct {
		name      string
		missing   []string
		wantCount int
	}{
		{
			name:      "nil returns empty",
			missing:   nil,
			wantCount: 0,
		},
		{
			name:      "empty returns empty",
			missing:   []string{},
			wantCount: 0,
		},
		{
			name:      "one missing path emits one diagnostic",
			missing:   []string{"db/migrate/001_backfill.sql"},
			wantCount: 1,
		},
		{
			name:      "multiple missing paths emit multiple diagnostics",
			missing:   []string{"db/migrate/001.sql", "db/migrate/002.sql"},
			wantCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := Mig003DataMigrationMissing(tt.missing)
			assertMig003Diags(t, diags, tt.wantCount, tt.missing)
		})
	}
}
