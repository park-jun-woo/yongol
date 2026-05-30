//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestWriteXsm27OptionB — writeXsm27OptionB Option B advice 라인 작성 검증

package ssac_statemachine

import (
	"strings"
	"testing"
)

func TestWriteXsm27OptionB(t *testing.T) {
	var b strings.Builder
	writeXsm27OptionB(&b)
	out := b.String()
	if !strings.Contains(out, "Option B (state-neutral)") || !strings.Contains(out, "@state-neutral") {
		t.Errorf("unexpected output: %s", out)
	}
}
