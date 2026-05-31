//ff:func feature=generate type=test control=iteration dimension=1
//ff:what TestZeroCov — 0% util 함수 (isCopiedExtension / isYongolManaged / mergeFieldlessOps / ResolveBackendType / WithMigration / appendChildNodeFormActions) 회귀
package generate

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestAppendChildNodeFormActions(t *testing.T) {
	seen := map[string]bool{}
	// action branch with fields
	res := appendChildNodeFormActions(nil, stmlparser.ChildNode{
		Kind:   "action",
		Action: &stmlparser.ActionBlock{OperationID: "CreateX", Fields: []stmlparser.FieldBind{{Name: "n"}}},
	}, seen)
	if len(res) != 1 || res[0].opID != "CreateX" {
		t.Fatalf("action branch = %+v", res)
	}
	// duplicate operationId is skipped
	res = appendChildNodeFormActions(res, stmlparser.ChildNode{
		Kind:   "action",
		Action: &stmlparser.ActionBlock{OperationID: "CreateX", Fields: []stmlparser.FieldBind{{Name: "n"}}},
	}, seen)
	if len(res) != 1 {
		t.Errorf("duplicate not skipped: %+v", res)
	}
	// recurse branches: fetch / state / static / each with nested action
	nested := stmlparser.ChildNode{
		Kind:   "action",
		Action: &stmlparser.ActionBlock{OperationID: "Nested", Fields: []stmlparser.FieldBind{{Name: "m"}}},
	}
	branches := []stmlparser.ChildNode{
		{Kind: "fetch", Fetch: &stmlparser.FetchBlock{Children: []stmlparser.ChildNode{nested}}},
		{Kind: "state", State: &stmlparser.StateBind{Children: []stmlparser.ChildNode{nested}}},
		{Kind: "static", Static: &stmlparser.StaticElement{Children: []stmlparser.ChildNode{nested}}},
		{Kind: "each", Each: &stmlparser.EachBlock{Children: []stmlparser.ChildNode{nested}}},
	}
	for _, br := range branches {
		out := appendChildNodeFormActions(nil, br, map[string]bool{})
		if len(out) != 1 || out[0].opID != "Nested" {
			t.Errorf("recurse branch %q = %+v", br.Kind, out)
		}
	}
	// unknown kind → no change
	out := appendChildNodeFormActions(nil, stmlparser.ChildNode{Kind: "bind"}, map[string]bool{})
	if len(out) != 0 {
		t.Errorf("bind kind should add nothing: %+v", out)
	}
}
