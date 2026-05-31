//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestHelpersZeroCov — head 판정 / native·pgtype·array·enum·jsonb·bytea binding 헬퍼 직접 커버
package types

import (
	"testing"
)

func TestNativeBindings_ZeroCov(t *testing.T) {
	// NOT NULL native forms.
	if b := nativeInteger(true, "0"); b.SqlcGoType != "int64" || b.Kind != KindNative {
		t.Errorf("nativeInteger NOT NULL = %+v", b)
	}
	// Nullable → pgtype.
	if b := nativeInteger(false, "0"); b.Kind != KindPgtype {
		t.Errorf("nativeInteger nullable = %+v", b)
	}
	if b := nativeString(true, "''"); b.SqlcGoType != "string" || b.Kind != KindNative {
		t.Errorf("nativeString NOT NULL = %+v", b)
	}
	if b := nativeString(false, "''"); b.Kind != KindPgtype {
		t.Errorf("nativeString nullable = %+v", b)
	}
	if b := nativeBoolean(true, "false"); b.SqlcGoType != "bool" || b.Kind != KindNative {
		t.Errorf("nativeBoolean NOT NULL = %+v", b)
	}
	if b := nativeBoolean(false, "false"); b.Kind != KindPgtype {
		t.Errorf("nativeBoolean nullable = %+v", b)
	}
}
