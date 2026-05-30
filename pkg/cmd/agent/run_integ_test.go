//ff:func feature=agent type=test control=sequence
//ff:what run_integ — agent.Run 메인 흐름 (scaffold-only / validateLoop) best-effort 커버리지

package agent

import (
	"bytes"
	"strings"
	"testing"
)

// TestRun_ScaffoldOnly_Integ drives Run through its scaffold-only branch:
// an empty specs dir → features.Load yields nil ff → scaffold is skipped →
// MaxRounds==0 → the "scaffold only" message is printed and Run returns nil.
func TestRun_ScaffoldOnly_Integ(t *testing.T) {
	var buf bytes.Buffer
	cfg := Config{SpecsDir: t.TempDir(), MaxRounds: 0}
	if err := Run(&buf, cfg); err != nil {
		t.Fatalf("Run scaffold-only: %v", err)
	}
	if !strings.Contains(buf.String(), "scaffold only") {
		t.Errorf("expected scaffold-only message, got: %q", buf.String())
	}
}

// TestRun_ValidateLoop_Integ drives Run into the validateLoop branch with an
// empty (valid) specs dir and no features.yaml: there is nothing to fix, so the
// loop terminates cleanly without any LLM call. This covers the MaxRounds>0
// path of Run end to end.
func TestRun_ValidateLoop_Integ(t *testing.T) {
	var buf bytes.Buffer
	cfg := Config{SpecsDir: t.TempDir(), MaxRounds: 1}
	if err := Run(&buf, cfg); err != nil {
		t.Fatalf("Run validate-loop: %v", err)
	}
}
