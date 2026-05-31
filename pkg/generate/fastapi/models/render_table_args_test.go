//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestModelsHelpers — fastapi models 패키지 순수 헬퍼(타입 매핑·PK·FK·기본값·table_args) 검증
package models

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestRenderTableArgs(t *testing.T) {
	// No indexes -> empty.
	if got := renderTableArgs(ddl.Table{}); got != "" {
		t.Errorf("expected empty for no indexes, got %q", got)
	}
	// Unique + non-unique index.
	tbl := ddl.Table{
		Indexes: []ddl.Index{
			{Name: "uq_email", Columns: []string{"email"}, IsUnique: true},
			{Name: "idx_name", Columns: []string{"first", "last"}, IsUnique: false},
		},
	}
	got := renderTableArgs(tbl)
	if !strings.Contains(got, `UniqueConstraint("email", name="uq_email")`) {
		t.Errorf("missing unique constraint: %q", got)
	}
	if !strings.Contains(got, `Index("idx_name", "first", "last")`) {
		t.Errorf("missing index: %q", got)
	}
	if !strings.HasPrefix(got, "    __table_args__ = (\n") || !strings.HasSuffix(got, "    )\n") {
		t.Errorf("malformed table_args wrapper: %q", got)
	}
}
