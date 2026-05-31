//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestZeroCov — 0% util/adapter 함수 (bindGoType / GoTypeOf / registry / family dispatch / pgtype) 회귀
package types

import (
	"testing"
)

func TestPgtypeInt2Int4(t *testing.T) {
	i2 := pgtypeInt2("0")
	if i2.SqlcGoType != "pgtype.Int2" || i2.Kind != KindPgtype {
		t.Errorf("pgtypeInt2 = %+v", i2)
	}
	i4 := pgtypeInt4("0")
	if i4.SqlcGoType != "pgtype.Int4" || i4.Kind != KindPgtype {
		t.Errorf("pgtypeInt4 = %+v", i4)
	}
}
