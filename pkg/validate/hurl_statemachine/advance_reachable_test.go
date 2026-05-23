//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-statemachine
//ff:what advanceReachable — op에 해당하는 전이의 To 상태를 reachable에 추가 검증

package hurl_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestAdvanceReachable(t *testing.T) {
	d := &statemachine.StateDiagram{
		Transitions: []statemachine.Transition{
			{From: "draft", To: "published", Event: "publish"},
			{From: "draft", To: "archived", Event: "archive"},
			{From: "published", To: "archived", Event: "archive"},
		},
	}

	cases := []struct {
		name     string
		op       string
		wantKeys []string
	}{
		{name: "publish_adds_published", op: "publish", wantKeys: []string{"published"}},
		{name: "archive_adds_archived", op: "archive", wantKeys: []string{"archived"}},
		{name: "unknown_op_adds_nothing", op: "delete", wantKeys: nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runAdvanceReachableCase(t, d, c.op, c.wantKeys)
		})
	}
}
