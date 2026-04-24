//ff:func feature=validate type=test-helper control=sequence topic=states
//ff:what runXsm27Case — XSM-27 단일 케이스 실행 + 기대 결과 검증

package ssac_statemachine

import (
	"strings"
	"testing"
)

// runXsm27Case executes one xsm27Case and asserts wantFire along with the
// required message / advice substrings.
func runXsm27Case(t *testing.T, tc xsm27Case) {
	t.Helper()
	fs := buildXsm27Fixture(tc)
	diags := xsm27StateIntentDeclaration(fs)
	got := len(diags) > 0
	if got != tc.wantFire {
		t.Fatalf("want fire=%v, got %d diagnostics: %+v", tc.wantFire, len(diags), diags)
	}
	if !tc.wantFire {
		return
	}
	d := diags[0]
	if !strings.Contains(d.Message, "[XSM-27]") {
		t.Errorf("expected [XSM-27] in message, got %q", d.Message)
	}
	if !strings.Contains(d.Advice, "Option A") || !strings.Contains(d.Advice, "Option B") {
		t.Errorf("advice must carry Option A and Option B, got %q", d.Advice)
	}
	if !strings.Contains(d.Advice, "@state-neutral") {
		t.Errorf("advice must mention @state-neutral, got %q", d.Advice)
	}
}
