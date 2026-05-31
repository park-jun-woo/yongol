//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestRenderSchema_ZeroCov(t *testing.T) {
	tables := []ddl.Table{
		{
			Name:        "users",
			ColumnOrder: []string{"id"},
			Columns:     map[string]ddl.Column{"id": {RawType: "BIGINT", NotNull: true}},
			PrimaryKey:  []string{"id"},
		},
		{
			Name:        "posts",
			ColumnOrder: []string{"id", "user_id"},
			Columns: map[string]ddl.Column{
				"id":      {RawType: "BIGINT", NotNull: true},
				"user_id": {RawType: "BIGINT", NotNull: true},
			},
			PrimaryKey:  []string{"id"},
			ForeignKeys: []ddl.ForeignKey{{Column: "user_id", RefTable: "users", RefColumn: "id"}},
		},
	}
	out, err := RenderSchema(tables)
	if err != nil {
		t.Fatalf("RenderSchema error: %v", err)
	}
	for _, want := range []string{"datasource db {", "generator client {", "model User {", "model Post {"} {
		if !strings.Contains(out, want) {
			t.Errorf("schema missing %q", want)
		}
	}
}
