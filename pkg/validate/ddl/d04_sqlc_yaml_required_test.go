//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestZeroCov — 0% DDL rule 함수 (d01/d03/d04/d05/d06/d07/d15/xdd61/parseSqlcYaml/Run) 회귀
package ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestD04SqlcYamlRequired(t *testing.T) {
	// no sqlc.yaml → ERROR
	fs := zcWriteSpecs(t, "", "CREATE TABLE t (id BIGINT);")
	if d := d04SqlcYamlRequired(fs); len(d) != 1 {
		t.Errorf("missing sqlc.yaml want 1 diag, got %d", len(d))
	}
	// present → none
	fs2 := zcWriteSpecs(t, zcGoodSqlc, "")
	if d := d04SqlcYamlRequired(fs2); len(d) != 0 {
		t.Errorf("present sqlc.yaml want 0 diag, got %d", len(d))
	}
	// DDL absent → none
	if d := d04SqlcYamlRequired(&yongol.Fullstack{}); len(d) != 0 {
		t.Errorf("absent DDL want 0 diag, got %d", len(d))
	}
}
