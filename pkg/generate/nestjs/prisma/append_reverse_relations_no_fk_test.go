//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestAppendReverseRelationsNoFK — TestAppendReverseRelations — 테이블 FK 들을 RefTable 키 아래 역관계로 추가 검증

package prisma

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestAppendReverseRelationsNoFK(t *testing.T) {
	rm := map[string][]reverseRelation{}
	appendReverseRelations(rm, ddl.Table{Name: "users"})
	if len(rm) != 0 {
		t.Errorf("rm = %v, want empty when table has no FK", rm)
	}
}
