//ff:func feature=validate type=test control=sequence topic=stml-statemachine
//ff:what collectPageStateConditions — state/fetch/static 래퍼 통과·each 미하강·기타 무시 검증

package stml_statemachine

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectPageStateConditions(t *testing.T) {
	t.Run("empty children returns nil", func(t *testing.T) {
		if got := collectPageStateConditions(nil); got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})

	t.Run("collects through state/fetch/static, ignores others", func(t *testing.T) {
		children := []stml.ChildNode{
			{Kind: "state", State: &stml.StateBind{Condition: "top"}},
			{Kind: "fetch", Fetch: &stml.FetchBlock{Children: []stml.ChildNode{
				{Kind: "state", State: &stml.StateBind{Condition: "inFetch"}},
			}}},
			{Kind: "static", Static: &stml.StaticElement{Children: []stml.ChildNode{
				{Kind: "state", State: &stml.StateBind{Condition: "inStatic"}},
			}}},
			{Kind: "bind", Bind: &stml.FieldBind{}}, // ignored
		}
		got := collectPageStateConditions(children)
		want := []string{"top", "inFetch", "inStatic"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
