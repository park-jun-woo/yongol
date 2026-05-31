//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderModel_ZeroCov — renderColumn / renderOneModel 을 DDL 테이블로 직접 호출
package models

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestRenderColumn_ZeroCov(t *testing.T) {
	var b strings.Builder
	// nullable column with default.
	col := ddl.Column{Name: "title", RawType: "TEXT", NotNull: false, HasDefault: true, DefaultLiteral: "x"}
	renderColumn(&b, col, "title", []string{"id"}, nil)
	if !strings.Contains(b.String(), "title:") || !strings.Contains(b.String(), "mapped_column") {
		t.Errorf("render = %q", b.String())
	}

	// primary key column.
	var b2 strings.Builder
	renderColumn(&b2, ddl.Column{Name: "id", RawType: "BIGINT", NotNull: true}, "id", []string{"id"}, nil)
	if !strings.Contains(b2.String(), "primary_key=True") {
		t.Errorf("pk render = %q", b2.String())
	}

	// foreign-key column.
	var b3 strings.Builder
	fks := []ddl.ForeignKey{{Column: "user_id", RefTable: "users", RefColumn: "id"}}
	renderColumn(&b3, ddl.Column{Name: "user_id", RawType: "BIGINT", NotNull: true}, "user_id", []string{"id"}, fks)
	if !strings.Contains(b3.String(), "ForeignKey") {
		t.Errorf("fk render = %q", b3.String())
	}
}
