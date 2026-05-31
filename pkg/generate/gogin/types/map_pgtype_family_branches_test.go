//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestZeroCov — 0% util/adapter 함수 (bindGoType / GoTypeOf / registry / family dispatch / pgtype) 회귀
package types

import (
	"testing"
)

func TestMapPgtypeFamily_Branches(t *testing.T) {
	heads := []string{"UUID", "NUMERIC", "DECIMAL", "TIMESTAMPTZ", "TIMESTAMP", "DATE", "INET", "CIDR", "INTERVAL", "JSONB", "JSON", "BYTEA"}
	for _, h := range heads {
		if _, ok := mapPgtypeFamily(h, true, ""); !ok {
			t.Errorf("%q should be pgtype family", h)
		}
	}
	if _, ok := mapPgtypeFamily("BIGINT", true, ""); ok {
		t.Errorf("BIGINT should not be pgtype family")
	}
}
