//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockEnvHelpers — envInt / envDuration / envStringList / envBool 헬퍼 (top-level)

package boot

import (
	"strings"
	"testing"
)

func TestBlockEnvHelpers(t *testing.T) {
	block := blockEnvHelpers()
	if block.Name != "env-helpers" {
		t.Errorf("name = %q, want env-helpers", block.Name)
	}
	// Payload lives entirely in Funcs (no main() body lines).
	if len(block.Lines) != 0 {
		t.Errorf("env helpers must emit no body lines, got %v", block.Lines)
	}
	funcs := strings.Join(block.Funcs, "\n")
	for _, must := range []string{
		"func envInt(",
		"func envInt64(",
		"func envDuration(",
		"func envStringList(",
		"func envBool(",
		"func envString(",
		"func envFloat(",
	} {
		if !strings.Contains(funcs, must) {
			t.Errorf("env helpers missing %q", must)
		}
	}
	imp := strings.Join(block.Imports, "\n")
	for _, must := range []string{`"strconv"`, `"time"`, `"strings"`, `"os"`} {
		if !strings.Contains(imp, must) {
			t.Errorf("env helpers missing import %q, got:\n%s", must, imp)
		}
	}
}
