//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what collectEachRowActionsChild — each 액션 채집 / state·link 컨테이너 재귀 / 무관 Kind 무시 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectEachRowActionsChild(t *testing.T) {
	t.Run("each contributes its row actions", func(t *testing.T) {
		var out []stml.ActionBlock
		collectEachRowActionsChild(stml.ChildNode{Kind: "each", Each: &stml.EachBlock{
			Field:   "items",
			Actions: []stml.ActionBlock{{OperationID: "DeleteBuilding"}},
		}}, &out)
		if len(out) != 1 || out[0].OperationID != "DeleteBuilding" {
			t.Errorf("out = %+v, want the each row action", out)
		}
	})

	t.Run("state and link containers recurse", func(t *testing.T) {
		var out []stml.ActionBlock
		inner := stml.ChildNode{Kind: "each", Each: &stml.EachBlock{Field: "items", Actions: []stml.ActionBlock{{OperationID: "DeleteUnit"}}}}
		collectEachRowActionsChild(stml.ChildNode{Kind: "state", State: &stml.StateBind{Children: []stml.ChildNode{inner}}}, &out)
		collectEachRowActionsChild(stml.ChildNode{Kind: "link", Link: &stml.LinkRef{TargetPage: "x", Children: []stml.ChildNode{inner}}}, &out)
		if len(out) != 2 {
			t.Errorf("out = %+v, want 2 nested row actions", out)
		}
	})

	t.Run("unrelated kinds are ignored", func(t *testing.T) {
		var out []stml.ActionBlock
		collectEachRowActionsChild(stml.ChildNode{Kind: "bind", Bind: &stml.FieldBind{}}, &out)
		if out != nil {
			t.Errorf("out = %+v, want nil for a bind node", out)
		}
	})
}
