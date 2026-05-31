//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderModel_ZeroCov — renderColumn / renderOneModel 을 DDL 테이블로 직접 호출
package models

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestRenderOneModel_ZeroCov(t *testing.T) {
	var b strings.Builder
	table := ddl.Table{
		Name:        "orders",
		ColumnOrder: []string{"id", "title", "missing"},
		Columns: map[string]ddl.Column{
			"id":    {Name: "id", RawType: "BIGINT", NotNull: true},
			"title": {Name: "title", RawType: "TEXT"},
			// "missing" intentionally absent from Columns → skipped branch.
		},
		PrimaryKey: []string{"id"},
	}
	if err := renderOneModel(&b, table); err != nil {
		t.Fatalf("renderOneModel: %v", err)
	}
	out := b.String()
	if !strings.Contains(out, "class Order(Base):") {
		t.Errorf("model class missing: %q", out)
	}
	if !strings.Contains(out, `__tablename__ = "orders"`) {
		t.Errorf("tablename missing: %q", out)
	}
}
