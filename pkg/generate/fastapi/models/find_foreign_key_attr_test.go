//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestModelsHelpers — fastapi models 패키지 순수 헬퍼(타입 매핑·PK·FK·기본값·table_args) 검증
package models

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestFindForeignKeyAttr(t *testing.T) {
	fks := []ddl.ForeignKey{
		{Column: "user_id", RefTable: "users", RefColumn: "id"},
	}
	if got := findForeignKeyAttr(fks, "user_id"); got != `ForeignKey("users.id")` {
		t.Errorf("unexpected fk attr: %q", got)
	}
	if got := findForeignKeyAttr(fks, "other"); got != "" {
		t.Errorf("expected empty for non-matching col, got %q", got)
	}
	if got := findForeignKeyAttr(nil, "user_id"); got != "" {
		t.Errorf("expected empty for nil fks, got %q", got)
	}
}
