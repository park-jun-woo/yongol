//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestSchemaPascalCase — camelCase/snake_case/PascalCase → PascalCase 변환 검증
package fastapi

import (
	"testing"
)

func TestSchemaPascalCase(t *testing.T) {
	cases := map[string]string{
		"":               "",
		"createWorkflow": "CreateWorkflow",
		"CreateWorkflow": "CreateWorkflow",
		"x":              "X",
	}
	for in, want := range cases {
		if got := schemaPascalCase(in); got != want {
			t.Errorf("schemaPascalCase(%q) = %q, want %q", in, got, want)
		}
	}
}
