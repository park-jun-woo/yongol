//ff:func feature=validate type=test control=sequence topic=statemachine-structural
//ff:what ST-1 positive 테스트 — states/*.md 파싱 실패 시 [ST-1] 진단 발화

package statemachine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	smparser "github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestST01ParseFiresOnInvalidMermaid(t *testing.T) {
	dir := t.TempDir()
	statesDir := filepath.Join(dir, "states")
	if err := os.MkdirAll(statesDir, 0o755); err != nil {
		t.Fatal(err)
	}
	// Write an .md file with no mermaid block — parser returns a diagnostic.
	bad := filepath.Join(statesDir, "order.md")
	if err := os.WriteFile(bad, []byte("# plain markdown, no mermaid\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := st01Parse(fs)
	if len(diags) == 0 {
		t.Fatalf("expected ST-1 diagnostic, got none")
	}
	if !strings.Contains(diags[0].Message, "[ST-1]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
}

// TestST01ParseSkipsWhenLoaded ensures the rule stays silent when parse already
// succeeded (fs.StateDiagrams non-empty → Run shortcut returns nil).
func TestST01ParseSkipsWhenLoaded(t *testing.T) {
	fs := &yongol.Fullstack{StateDiagrams: []*smparser.StateDiagram{{}}}
	diags := st01Parse(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics when diagrams pre-loaded, got %d", len(diags))
	}
}
