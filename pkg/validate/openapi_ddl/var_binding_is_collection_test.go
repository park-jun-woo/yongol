//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what varBindingIsCollection — var 의 Result 바인딩이 Wrapper/[] 컬렉션인지 판별, 미존재는 false

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestVarBindingIsCollection(t *testing.T) {
	fn := &ssac.ServiceFunc{Sequences: []ssac.Sequence{
		{Type: "get", Result: &ssac.Result{Var: "page", Type: "Gig", Wrapper: "Page"}},
		{Type: "get", Result: &ssac.Result{Var: "slice", Type: "[]Gig"}},
		{Type: "get", Result: &ssac.Result{Var: "one", Type: "Gig"}},
		{Type: "empty", Result: nil},
	}}
	if !varBindingIsCollection(fn, "page") {
		t.Error("page (Wrapper) should be collection")
	}
	if !varBindingIsCollection(fn, "slice") {
		t.Error("slice ([]) should be collection")
	}
	if varBindingIsCollection(fn, "one") {
		t.Error("one (single) should not be collection")
	}
	if varBindingIsCollection(fn, "missing") {
		t.Error("missing var should not be collection")
	}
}
