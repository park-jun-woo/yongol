//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestWriteXsm27SelfLoopHint — writeXsm27SelfLoopHint self-loop 힌트 작성 검증

package ssac_statemachine

import (
	"strings"
	"testing"
)

func TestWriteXsm27SelfLoopHint(t *testing.T) {
	var b strings.Builder
	writeXsm27SelfLoopHint(&b, "CancelOrder", "order", "draft")
	out := b.String()
	for _, want := range []string{
		"If CancelOrder is not declared as a transition in states/order.md",
		"draft --> draft: CancelOrder",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q\ngot: %s", want, out)
		}
	}
}
