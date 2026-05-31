//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestZeroCov — 0% util/adapter 함수 (bindGoType / GoTypeOf / registry / family dispatch / pgtype) 회귀
package types

import (
	"testing"
)

func TestBindingKindString(t *testing.T) {
	cases := []struct {
		k    BindingKind
		want string
	}{
		{KindNative, "Native"},
		{KindPointer, "Pointer"},
		{KindPgtype, "Pgtype"},
		{KindJSONB, "JSONB"},
		{KindBytea, "Bytea"},
		{KindArray, "Array"},
		{KindEnum, "Enum"},
		{KindUnsupported, "Unsupported"},
		{BindingKind(99), "Unknown"},
	}
	for _, c := range cases {
		if got := c.k.String(); got != c.want {
			t.Errorf("BindingKind(%d).String() = %q, want %q", c.k, got, c.want)
		}
	}
}
