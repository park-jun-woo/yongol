//ff:func feature=generate type=test control=iteration dimension=1
//ff:what TestGenerateBatch_ZeroCov — pkg/generate 소형 순수/IO 헬퍼 분기 커버
package generate

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestAppendFieldlessFromChild_ZeroCov(t *testing.T) {
	result := map[string]bool{}
	// action with no fields → recorded
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "action", Action: &stmlparser.ActionBlock{OperationID: "Op1"}}, result)
	if !result["Op1"] {
		t.Error("expected Op1 recorded")
	}
	// action with fields → not recorded
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "action", Action: &stmlparser.ActionBlock{OperationID: "Op2", Fields: []stmlparser.FieldBind{{Name: "x"}}}}, result)
	if result["Op2"] {
		t.Error("Op2 has fields, should not be recorded")
	}
	// nil action
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "action"}, result)
	// fetch / state / static / each recurse
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "fetch", Fetch: &stmlparser.FetchBlock{Children: []stmlparser.ChildNode{{Kind: "action", Action: &stmlparser.ActionBlock{OperationID: "F"}}}}}, result)
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "state", State: &stmlparser.StateBind{Children: []stmlparser.ChildNode{{Kind: "action", Action: &stmlparser.ActionBlock{OperationID: "S"}}}}}, result)
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "static", Static: &stmlparser.StaticElement{Children: []stmlparser.ChildNode{{Kind: "action", Action: &stmlparser.ActionBlock{OperationID: "T"}}}}}, result)
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "each", Each: &stmlparser.EachBlock{Children: []stmlparser.ChildNode{{Kind: "action", Action: &stmlparser.ActionBlock{OperationID: "E"}}}}}, result)
	for _, k := range []string{"F", "S", "T", "E"} {
		if !result[k] {
			t.Errorf("expected %q recorded via recursion", k)
		}
	}
	// nil pointers in recursive kinds → no panic
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "fetch"}, result)
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "state"}, result)
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "static"}, result)
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "each"}, result)
	// unknown kind
	appendFieldlessFromChild(stmlparser.ChildNode{Kind: "bind"}, result)
}
