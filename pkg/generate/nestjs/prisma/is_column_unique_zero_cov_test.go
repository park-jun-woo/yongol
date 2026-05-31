//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsColumnUnique_ZeroCov(t *testing.T) {
	idxs := []ddl.Index{
		{Name: "u1", Columns: []string{"email"}, IsUnique: true},
		{Name: "u2", Columns: []string{"a", "b"}, IsUnique: true},
		{Name: "n1", Columns: []string{"name"}, IsUnique: false},
	}
	if !isColumnUnique("email", idxs) {
		t.Error("expected email unique")
	}
	if isColumnUnique("a", idxs) {
		t.Error("composite index should not count as single-column unique")
	}
	if isColumnUnique("name", idxs) {
		t.Error("non-unique index should not count")
	}
}
