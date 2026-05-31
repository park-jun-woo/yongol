//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildTransitionMap_ZeroCov — 초기 전이 제외 + 중첩 맵 구성
package state

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestBuildTransitionMap_ZeroCov(t *testing.T) {
	d := &statemachine.StateDiagram{
		Transitions: []statemachine.Transition{
			{From: "[*]", To: "draft", Event: "create"},
			{From: "draft", To: "review", Event: "submit"},
			{From: "draft", To: "archived", Event: "archive"},
			{From: "review", To: "published", Event: "approve"},
		},
	}
	m := buildTransitionMap(d)
	if _, ok := m["[*]"]; ok {
		t.Error("initial transition should be excluded")
	}
	if m["draft"]["submit"] != "review" {
		t.Errorf("draft/submit = %q", m["draft"]["submit"])
	}
	if len(m["draft"]) != 2 {
		t.Errorf("draft events = %d", len(m["draft"]))
	}
}
