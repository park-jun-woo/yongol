//ff:func feature=validate type=test control=iteration dimension=1
//ff:what TestAllSteps_ZeroCov — allSteps() 고정 step 목록을 이름으로 직접 호출
package validate

import (
	"testing"
)

func TestAllSteps_ZeroCov(t *testing.T) {
	steps := allSteps()
	if len(steps) == 0 {
		t.Fatal("allSteps() returned no steps")
	}
	// First step is always the "init" gate.
	if steps[0].Name != "init" {
		t.Errorf("first step = %q, want init", steps[0].Name)
	}
	// Every step must carry a runnable function.
	for _, s := range steps {
		if s.Name == "" {
			t.Errorf("step with empty name: %+v", s)
		}
		if s.Run == nil {
			t.Errorf("step %q has nil Run", s.Name)
		}
	}
}
