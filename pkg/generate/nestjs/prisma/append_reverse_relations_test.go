//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestAppendReverseRelations — TestAppendReverseRelations — 테이블 FK 들을 RefTable 키 아래 역관계로 추가 검증

package prisma

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestAppendReverseRelations(t *testing.T) {
	rm := map[string][]reverseRelation{}
	table := ddl.Table{
		Name: "posts",
		ForeignKeys: []ddl.ForeignKey{
			{Column: "user_id", RefTable: "users", RefColumn: "id"},
			{Column: "org_id", RefTable: "orgs", RefColumn: "id"},
		},
	}

	appendReverseRelations(rm, table)

	for _, ref := range []string{"users", "orgs"} {
		revs := rm[ref]
		if len(revs) != 1 {
			t.Fatalf("rm[%q] len = %d, want 1", ref, len(revs))
		}
		// ModelName = PascalCase(singular("posts")) = "Post"
		if revs[0].ModelName != "Post" {
			t.Errorf("rm[%q].ModelName = %q, want Post", ref, revs[0].ModelName)
		}
		// FieldName = source table name ("posts")
		if revs[0].FieldName != "posts" {
			t.Errorf("rm[%q].FieldName = %q, want posts", ref, revs[0].FieldName)
		}
	}
}
