//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestAssembleRego — 헤더 고정 + non-empty 블록 추가 + clean 후 empty 블록 skip 검증

package agent

import (
	"strings"
	"testing"
)

func TestAssembleRego(t *testing.T) {
	const header = "package authz\n\nimport rego.v1\n\ndefault allow := false\n"

	t.Run("NoBlocks", func(t *testing.T) {
		got := assembleRego(nil)
		if got != header {
			t.Errorf("empty blocks should produce just the header.\n got: %q\nwant: %q", got, header)
		}
	})

	t.Run("RealBlockAppended", func(t *testing.T) {
		block := "allow if {\n  input.action == \"read\"\n}"
		got := assembleRego([]string{block})
		if !strings.HasPrefix(got, header) {
			t.Errorf("output must start with header, got: %q", got)
		}
		if !strings.Contains(got, "allow if {") {
			t.Errorf("expected rule body appended, got: %q", got)
		}
	})

	t.Run("CleanedEmptyBlockSkipped", func(t *testing.T) {
		// A block containing only duplicate package/import/default lines cleans
		// to "" and must be skipped (no extra body beyond the header).
		dupOnly := "package authz\nimport rego.v1\ndefault allow := false"
		got := assembleRego([]string{dupOnly})
		if got != header {
			t.Errorf("cleaned-empty block should be skipped.\n got: %q\nwant: %q", got, header)
		}
	})

	t.Run("MixedBlocks", func(t *testing.T) {
		blocks := []string{
			"package authz", // cleans to empty → skipped
			"allow if { input.role == \"admin\" }",
		}
		got := assembleRego(blocks)
		if !strings.Contains(got, "input.role == \"admin\"") {
			t.Errorf("expected the non-empty block, got: %q", got)
		}
		if strings.Count(got, "package authz") != 1 {
			t.Errorf("expected exactly one package line (header only), got: %q", got)
		}
	})
}
