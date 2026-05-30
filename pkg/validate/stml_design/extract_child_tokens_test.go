//ff:func feature=validate type=test control=selection topic=stml-design
//ff:what TestExtractChildTokens — extractChildTokens ChildNode 종류별 분기 검증

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestExtractChildTokens(t *testing.T) {
	cases := []struct {
		name string
		node stml.ChildNode
	}{
		{name: "static", node: stml.ChildNode{Kind: "static", Static: &stml.StaticElement{ClassName: "bg-primary"}}},
		{name: "fetch", node: stml.ChildNode{Kind: "fetch", Fetch: &stml.FetchBlock{ClassName: "bg-primary"}}},
		{name: "action", node: stml.ChildNode{Kind: "action", Action: &stml.ActionBlock{ClassName: "bg-primary"}}},
		{name: "each", node: stml.ChildNode{Kind: "each", Each: &stml.EachBlock{ClassName: "bg-primary"}}},
		{name: "component", node: stml.ChildNode{Kind: "component", Component: &stml.ComponentRef{Name: "DatePicker", ClassName: "bg-primary"}}},
		{name: "bind", node: stml.ChildNode{Kind: "bind", Bind: &stml.FieldBind{ClassName: "bg-primary"}}},
		{name: "static nil ptr", node: stml.ChildNode{Kind: "static"}},
		{name: "fetch nil ptr", node: stml.ChildNode{Kind: "fetch"}},
		{name: "action nil ptr", node: stml.ChildNode{Kind: "action"}},
		{name: "each nil ptr", node: stml.ChildNode{Kind: "each"}},
		{name: "component nil ptr", node: stml.ChildNode{Kind: "component"}},
		{name: "bind nil ptr", node: stml.ChildNode{Kind: "bind"}},
		{name: "unknown kind", node: stml.ChildNode{Kind: "mystery"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var out pageTokenRefs
			// Must not panic for any branch.
			extractChildTokens(tc.node, "p.stml", &out)
		})
	}
}
