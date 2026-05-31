//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestCheckChildClass — ChildNode Kind별 디스패치 + nil 포인터/미지 Kind 분기 검증
package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCheckChildClass(t *testing.T) {
	// Non-nil pointers for every Kind, each carrying a class so a diagnostic surfaces.
	nonNil := []stml.ChildNode{
		{Kind: "static", Static: &stml.StaticElement{Tag: "div", ClassName: "p-2"}},
		{Kind: "fetch", Fetch: &stml.FetchBlock{OperationID: "List", ClassName: "m-1"}},
		{Kind: "action", Action: &stml.ActionBlock{OperationID: "Create", ClassName: "bg-x"}},
		{Kind: "each", Each: &stml.EachBlock{ClassName: "gap-2"}},
		{Kind: "state", State: &stml.StateBind{Condition: "empty", ClassName: "text-x"}},
		{Kind: "component", Component: &stml.ComponentRef{Name: "Modal", ClassName: "shadow"}},
		{Kind: "bind", Bind: &stml.FieldBind{Name: "f", ClassName: "w-1"}},
	}
	for _, cn := range nonNil {
		if got := checkChildClass(cn, "p.stml"); len(got) == 0 {
			t.Errorf("kind %q with class expected a diagnostic, got none", cn.Kind)
		}
	}

	// Same Kinds but nil sub-pointers → no diagnostics (nil guard branches).
	nilPtrs := []stml.ChildNode{
		{Kind: "static"},
		{Kind: "fetch"},
		{Kind: "action"},
		{Kind: "each"},
		{Kind: "state"},
		{Kind: "component"},
		{Kind: "bind"},
		{Kind: "unknown"}, // default branch
	}
	for _, cn := range nilPtrs {
		if got := checkChildClass(cn, "p.stml"); got != nil {
			t.Errorf("kind %q nil ptr expected nil, got %+v", cn.Kind, got)
		}
	}
}
