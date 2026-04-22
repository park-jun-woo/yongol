//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what requiresAuth 매트릭스 단위 테스트
package boot

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestRequiresAuth_Matrix(t *testing.T) {
	docWithGlobal := &openapi3.T{
		Security: openapi3.SecurityRequirements{
			{"bearerAuth": []string{}},
		},
	}
	docBare := &openapi3.T{}

	opNil := &openapi3.Operation{}
	opEmpty := &openapi3.Operation{Security: &openapi3.SecurityRequirements{}}
	opSet := &openapi3.Operation{
		Security: &openapi3.SecurityRequirements{
			{"bearerAuth": []string{}},
		},
	}

	cases := []struct {
		name string
		op   *openapi3.Operation
		doc  *openapi3.T
		want bool
	}{
		{"nil op nil doc", opNil, docBare, false},
		{"nil op global", opNil, docWithGlobal, true},
		{"empty op global (opt-out)", opEmpty, docWithGlobal, false},
		{"explicit op (override)", opSet, docBare, true},
		{"nil op nil doc literal", opNil, nil, false},
	}
	for _, c := range cases {
		if got := requiresAuth(c.op, c.doc); got != c.want {
			t.Errorf("%s: got %v want %v", c.name, got, c.want)
		}
	}

	if requiresAuth(nil, docWithGlobal) != false {
		t.Errorf("nil op should be false")
	}
}
