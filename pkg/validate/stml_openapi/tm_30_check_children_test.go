//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm30CheckChildren — 슬라이스 순회로 진단이 누적됨을 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM30CheckChildren(t *testing.T) {
	a := &stml.ActionBlock{
		OperationID: "DeletePhoto",
		Params:      []stml.ParamBind{{Name: "photoId", Source: "item.id"}},
	}
	b := &stml.ActionBlock{
		OperationID: "StarPhoto",
		Params:      []stml.ParamBind{{Name: "photoId", Source: "item.id"}},
	}
	var out []diagnostic.Diagnostic
	children := []stml.ChildNode{
		{Kind: "action", Action: a},
		{Kind: "action", Action: b},
	}
	// both actions sit outside any each → two diagnostics accumulate
	tm30CheckChildren(children, "p.html", "", nil, false, nil, &out)
	if len(out) != 2 {
		t.Errorf("expected 2 diags, got %d: %v", len(out), out)
	}
}
