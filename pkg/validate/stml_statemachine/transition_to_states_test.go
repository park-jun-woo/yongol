//ff:func feature=validate type=test control=sequence topic=stml-statemachine
//ff:what transitionToStates — 이벤트 라벨 전이의 도착 상태 집합 / 미매칭 빈 결과 검증

package stml_statemachine

import (
	"reflect"
	"testing"
)

func TestTransitionToStates(t *testing.T) {
	d := workflowDiagram()

	t.Run("returns arrival states for the event", func(t *testing.T) {
		got := transitionToStates(d, "ActivateWorkflow")
		if !reflect.DeepEqual(got, []string{"active"}) {
			t.Errorf("expected [active], got %v", got)
		}
	})

	t.Run("unknown event returns empty", func(t *testing.T) {
		if got := transitionToStates(d, "CreateWorkflow"); len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})
}
