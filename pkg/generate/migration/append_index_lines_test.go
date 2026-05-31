//ff:func feature=migration type=test control=iteration dimension=1
//ff:what render_helpers_unit_test — appendColumnLines/PK/FK/Check/Index/Sentinel + renderTable + CanonicalSQL 단위 테스트
package migration

import (
	"strings"
	"testing"
)

func TestAppendIndexLines(t *testing.T) {
	tbl := &Table{
		Name: "users",
		Indexes: []*Index{
			{Name: "idx_b", Columns: []string{"b"}},
			{Name: "idx_a", Columns: []string{"a"}, Unique: true},
			{Name: "idx_gin", Columns: []string{"doc"}, Method: "gin"},
			{Name: "idx_w", Columns: []string{"a"}, Where: "a IS NOT NULL"},
		},
	}
	var b strings.Builder
	appendIndexLines(&b, tbl)
	got := b.String()
	lines := strings.Split(strings.TrimRight(got, "\n"), "\n")
	// name-sorted: idx_a, idx_b, idx_gin, idx_w
	want := []string{
		"CREATE UNIQUE INDEX idx_a ON users (a);",
		"CREATE INDEX idx_b ON users (b);",
		"CREATE INDEX idx_gin ON users USING gin (doc);",
		"CREATE INDEX idx_w ON users (a) WHERE a IS NOT NULL;",
	}
	if len(lines) != len(want) {
		t.Fatalf("got %d lines, want %d: %q", len(lines), len(want), got)
	}
	for i := range want {
		if lines[i] != want[i] {
			t.Errorf("line %d = %q, want %q", i, lines[i], want[i])
		}
	}
}
