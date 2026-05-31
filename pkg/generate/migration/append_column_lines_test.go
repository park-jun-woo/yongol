//ff:func feature=migration type=test control=sequence
//ff:what render_helpers_unit_test — appendColumnLines/PK/FK/Check/Index/Sentinel + renderTable + CanonicalSQL 단위 테스트
package migration

import (
	"strings"
	"testing"
)

func TestAppendColumnLines(t *testing.T) {
	var b strings.Builder
	appendColumnLines(&b, []*Column{
		{Name: "id", Type: CanonicalType{Base: "INTEGER"}, Nullable: false},
		{Name: "name", Type: CanonicalType{Base: "TEXT"}, Nullable: true},
	})
	got := b.String()
	want := "\n    id INTEGER NOT NULL,\n    name TEXT"
	if got != want {
		t.Errorf("appendColumnLines = %q, want %q", got, want)
	}

	var empty strings.Builder
	appendColumnLines(&empty, nil)
	if empty.String() != "" {
		t.Errorf("empty cols should produce no output, got %q", empty.String())
	}
}
