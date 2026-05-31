//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestRenderModel_ZeroCov(t *testing.T) {
	var b strings.Builder
	table := ddl.Table{
		Name:        "posts",
		ColumnOrder: []string{"id", "user_id", "title", "missing"},
		Columns: map[string]ddl.Column{
			"id":      {RawType: "BIGINT", NotNull: true},
			"user_id": {RawType: "BIGINT", NotNull: true},
			"title":   {RawType: "TEXT"},
		},
		PrimaryKey:  []string{"id"},
		ForeignKeys: []ddl.ForeignKey{{Column: "user_id", RefTable: "users", RefColumn: "id"}},
		Indexes: []ddl.Index{
			{Columns: []string{"title"}, IsUnique: false},
			{Columns: []string{"user_id", "title"}, IsUnique: true},
		},
	}
	revRels := []reverseRelation{{FieldName: "comments", ModelName: "Comment"}}
	if err := renderModel(&b, table, revRels); err != nil {
		t.Fatalf("renderModel error: %v", err)
	}
	out := b.String()
	for _, want := range []string{"model Post {", "user", "User @relation", "comments", "Comment[]", "@@index([title])", "@@unique([user_id, title])", `@@map("posts")`} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q\n%s", want, out)
		}
	}
}
