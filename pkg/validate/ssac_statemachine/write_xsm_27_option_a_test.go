//ff:func feature=validate type=test control=iteration dimension=1 topic=states
//ff:what TestWriteXsm27OptionA — writeXsm27OptionA Option A advice 라인 작성 검증
package ssac_statemachine

import (
	"strings"
	"testing"
)

func TestWriteXsm27OptionA(t *testing.T) {
	var b strings.Builder
	target := &statefulTarget{StateColumn: "status"}
	writeXsm27OptionA(&b, "CancelOrder", target, "order", "order")
	out := b.String()
	for _, want := range []string{
		"Option A (state-dependent)",
		"func CancelOrder",
		"@state order {Status: order.Status}",
		"\"CancelOrder\"",
		"409",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q\ngot: %s", want, out)
		}
	}
}
