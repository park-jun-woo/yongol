//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestHelpersZeroCov — head 판정 / native·pgtype·array·enum·jsonb·bytea binding 헬퍼 직접 커버
package types

import (
	"testing"
)

func TestNativeFloatWithHead_ZeroCov(t *testing.T) {
	if b := nativeFloatWithHead("REAL", true, "0"); b.SqlcGoType != "float64" || b.Kind != KindNative {
		t.Errorf("float NOT NULL = %+v", b)
	}
	if b := nativeFloatWithHead("REAL", false, "0"); b.SqlcGoType != "pgtype.Float4" {
		t.Errorf("REAL nullable should be Float4: %+v", b)
	}
	if b := nativeFloatWithHead("FLOAT4", false, "0"); b.SqlcGoType != "pgtype.Float4" {
		t.Errorf("FLOAT4 nullable should be Float4: %+v", b)
	}
	if b := nativeFloatWithHead("FLOAT8", false, "0"); b.SqlcGoType != "pgtype.Float8" {
		t.Errorf("FLOAT8 nullable should be Float8: %+v", b)
	}
}
