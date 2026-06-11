//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what CollectChildActions — fetch/state/static/each 중첩 트리에서 ActionBlock DOM 순서 수집 검증

package stml

import "testing"

func TestCollectChildActions(t *testing.T) {
	nodes := []ChildNode{
		{Kind: "action", Action: &ActionBlock{OperationID: "TopAction"}},
		{Kind: "fetch", Fetch: &FetchBlock{Children: []ChildNode{
			{Kind: "state", State: &StateBind{Children: []ChildNode{
				{Kind: "action", Action: &ActionBlock{OperationID: "StateAction"}},
			}}},
			{Kind: "each", Each: &EachBlock{Children: []ChildNode{
				{Kind: "action", Action: &ActionBlock{OperationID: "EachAction"}},
			}}},
		}}},
		{Kind: "static", Static: &StaticElement{Children: []ChildNode{
			{Kind: "action", Action: &ActionBlock{OperationID: "StaticAction"}},
		}}},
		{Kind: "bind", Bind: &FieldBind{}},
	}
	got := CollectChildActions(nodes)
	want := []string{"TopAction", "StateAction", "EachAction", "StaticAction"}
	if len(got) != len(want) {
		t.Fatalf("got %d actions, want %d", len(got), len(want))
	}
	for i, w := range want {
		if got[i].OperationID != w {
			t.Errorf("got[%d] = %q, want %q", i, got[i].OperationID, w)
		}
	}
}
