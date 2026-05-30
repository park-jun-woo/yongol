//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestCheckFetchClass — FetchBlock 자체/Binds/Eaches/States/Components/Children/Nested class 검사

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCheckFetchClass(t *testing.T) {
	// Minimal block → no diagnostics.
	if got := checkFetchClass(stml.FetchBlock{OperationID: "List"}, "p.stml"); got != nil {
		t.Fatalf("empty fetch: expected nil, got %+v", got)
	}

	fb := stml.FetchBlock{
		OperationID:   "List",
		ClassName:     "section",
		Binds:         []stml.FieldBind{{Name: "t", ClassName: "b1"}},
		Eaches:        []stml.EachBlock{{Field: "items", ClassName: "e1"}},
		States:        []stml.StateBind{{Condition: "empty", ClassName: "s1"}},
		Components:    []stml.ComponentRef{{Name: "C", ClassName: "c1"}},
		Children:      []stml.ChildNode{{Kind: "bind", Bind: &stml.FieldBind{Name: "ch", ClassName: "ch1"}}},
		NestedFetches: []stml.FetchBlock{{OperationID: "Nested", ClassName: "n1"}},
	}
	diags := checkFetchClass(fb, "p.stml")
	// fetch(1) + bind(1) + each(1) + state(1) + component(1) + child(1) + nested(1) = 7.
	if len(diags) != 7 {
		t.Fatalf("expected 7 diagnostics, got %d: %+v", len(diags), diags)
	}
}
