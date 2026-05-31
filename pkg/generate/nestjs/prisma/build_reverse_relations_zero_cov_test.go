//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestBuildReverseRelations_ZeroCov(t *testing.T) {
	tables := []ddl.Table{
		{Name: "users"},
		{Name: "posts", ForeignKeys: []ddl.ForeignKey{{Column: "user_id", RefTable: "users", RefColumn: "id"}}},
	}
	rm := buildReverseRelations(tables)
	revs := rm["users"]
	if len(revs) != 1 || revs[0].ModelName != "Post" || revs[0].FieldName != "posts" {
		t.Errorf("unexpected reverse relations: %+v", revs)
	}
}
