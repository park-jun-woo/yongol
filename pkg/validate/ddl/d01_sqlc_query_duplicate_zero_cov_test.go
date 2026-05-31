//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestZeroCov — 0% DDL rule 함수 (d01/d03/d04/d05/d06/d07/d15/xdd61/parseSqlcYaml/Run) 회귀
package ddl

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestD01SqlcQueryDuplicate_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "GetUser", File: "a.sql", Line: 1},
			{Name: "GetUser", File: "b.sql", Line: 2},
		},
	}
	d := d01SqlcQueryDuplicate(fs)
	if len(d) < 1 {
		t.Errorf("duplicate want diags, got %d", len(d))
	}
	// no queries → none
	if d := d01SqlcQueryDuplicate(&yongol.Fullstack{}); len(d) != 0 {
		t.Errorf("empty want 0, got %d", len(d))
	}
	// unique → none
	fs2 := &yongol.Fullstack{SQLcQueries: []sqlcparser.QuerySpec{{Name: "A"}, {Name: "B"}}}
	if d := d01SqlcQueryDuplicate(fs2); len(d) != 0 {
		t.Errorf("unique want 0, got %d", len(d))
	}
}
