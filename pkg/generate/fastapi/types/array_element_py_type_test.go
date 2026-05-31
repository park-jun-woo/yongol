//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestArrayElementPyType — 배열 element head → Python/SA 타입 매핑 (4개 family + 미지원)
package types

import (
	"testing"
)

func TestArrayElementPyType(t *testing.T) {
	cases := []struct {
		name   string
		head   string
		wantSA string
		wantPy string
		wantOk bool
	}{
		{"Integer", "BIGINT", "Integer", "int", true},
		{"Float", "REAL", "Float", "float", true},
		{"String", "TEXT", "String", "str", true},
		{"Boolean", "BOOLEAN", "Boolean", "bool", true},
		{"Unsupported", "UUID", "", "", false},
		{"Unknown", "MADEUPTYPE", "", "", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, ok := arrayElementPyType(c.head)
			if ok != c.wantOk {
				t.Fatalf("ok = %v, want %v", ok, c.wantOk)
			}
			if got.sa != c.wantSA || got.py != c.wantPy {
				t.Errorf("got {sa:%q py:%q}, want {sa:%q py:%q}", got.sa, got.py, c.wantSA, c.wantPy)
			}
		})
	}
}
