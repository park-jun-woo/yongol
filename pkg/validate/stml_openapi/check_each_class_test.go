//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestCheckEachClass — EachBlock 자체/Item/Binds/States/Components/Children class 검사

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCheckEachClass(t *testing.T) {
	// Minimal block: empty ItemClassName branch, no descendants.
	if got := checkEachClass(stml.EachBlock{Field: "items"}, "p.stml"); got != nil {
		t.Fatalf("empty each: expected nil, got %+v", got)
	}

	eb := stml.EachBlock{
		Field:         "items",
		ClassName:     "grid",
		ItemClassName: "card",
		Binds:         []stml.FieldBind{{Name: "title", ClassName: "text-lg"}},
		States:        []stml.StateBind{{Condition: "empty", ClassName: "muted"}},
		Components:    []stml.ComponentRef{{Name: "Badge", ClassName: "rounded"}},
		Children:      []stml.ChildNode{{Kind: "bind", Bind: &stml.FieldBind{Name: "c", ClassName: "p-1"}}},
	}
	diags := checkEachClass(eb, "p.stml")
	// each class(1) + item class(1) + bind(1) + state(1) + component(1) + child(1) = 6.
	if len(diags) != 6 {
		t.Fatalf("expected 6 diagnostics, got %d: %+v", len(diags), diags)
	}
}
