//ff:func feature=agent type=test control=sequence
//ff:what run_integ — agent.Run 메인 흐름 (scaffold-only / validateLoop) best-effort 커버리지
package agent

import (
	"bytes"
	"testing"
)

func TestRun_ValidateLoop_Integ(t *testing.T) {
	var buf bytes.Buffer
	cfg := Config{SpecsDir: t.TempDir(), MaxRounds: 1}
	if err := Run(&buf, cfg); err != nil {
		t.Fatalf("Run validate-loop: %v", err)
	}
}
