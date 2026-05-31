//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestHelpersZeroCov — head 판정 / native·pgtype·array·enum·jsonb·bytea binding 헬퍼 직접 커버
package types

import (
	"testing"
)

func TestArrayBindings_ZeroCov(t *testing.T) {
	// arrayElementGoType across all four supported families + unsupported.
	cases := []struct {
		head string
		want string
		ok   bool
	}{
		{"BIGINT", "int64", true},
		{"REAL", "float64", true},
		{"TEXT", "string", true},
		{"BOOLEAN", "bool", true},
		{"UUID", "", false},
	}
	for _, c := range cases {
		got, ok := arrayElementGoType(c.head)
		if got != c.want || ok != c.ok {
			t.Errorf("arrayElementGoType(%q) = (%q,%v), want (%q,%v)", c.head, got, ok, c.want, c.ok)
		}
	}
	// arrayBinding supported → KindArray.
	if b := arrayBinding("TEXT", "'{}'"); b.Kind != KindArray || b.SqlcGoType != "[]string" {
		t.Errorf("arrayBinding TEXT = %+v", b)
	}
	// arrayBinding unsupported element → KindUnsupported.
	if b := arrayBinding("UUID", "'{}'"); b.Kind != KindUnsupported || b.Supported {
		t.Errorf("arrayBinding UUID should be unsupported: %+v", b)
	}
	// composeArrayBinding directly: supported false branch.
	if b := composeArrayBinding("", false, "UUID", "'{}'"); b.Supported {
		t.Errorf("composeArrayBinding unsupported should not be Supported")
	}
	if b := composeArrayBinding("int64", true, "BIGINT", "'{}'"); b.SqlcGoType != "[]int64" {
		t.Errorf("composeArrayBinding BIGINT = %+v", b)
	}
}
