//ff:func feature=agent type=test control=sequence
//ff:what run_integ — agent.Run 메인 흐름 (scaffold-only / validateLoop) best-effort 커버리지
package agent

import (
	"bytes"
	"strings"
	"testing"
)

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
