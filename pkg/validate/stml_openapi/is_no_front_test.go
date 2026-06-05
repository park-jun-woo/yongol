//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestIsNoFront — no-front 태그 유무 / nil op 판정

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestIsNoFront(t *testing.T) {
	cases := []struct {
		name string
		op   *openapi3.Operation
		want bool
	}{
		{"nil op", nil, false},
		{"no tags", &openapi3.Operation{}, false},
		{"other tag", &openapi3.Operation{Tags: []string{"users"}}, false},
		{"no-front tag", &openapi3.Operation{Tags: []string{"users", "no-front"}}, true},
	}

	for _, c := range cases {
		if got := isNoFront(c.op); got != c.want {
			t.Errorf("%s: isNoFront = %v, want %v", c.name, got, c.want)
		}
	}
}
