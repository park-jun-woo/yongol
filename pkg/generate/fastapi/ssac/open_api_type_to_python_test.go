//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestWriteHTTPHandlerBranches — writeHTTPHandler 미커버 분기(필수 쿼리/pre-auth/full)
package ssac

import (
	"testing"
)

func TestOpenAPITypeToPython(t *testing.T) {
	cases := map[string]string{
		"integer": "int",
		"number":  "float",
		"boolean": "bool",
		"string":  "str",
		"unknown": "str",
	}
	for in, want := range cases {
		if got := openAPITypeToPython(in); got != want {
			t.Errorf("openAPITypeToPython(%q) = %q, want %q", in, got, want)
		}
	}
}
