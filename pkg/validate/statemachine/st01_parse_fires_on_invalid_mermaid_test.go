//ff:func feature=validate type=test control=sequence topic=statemachine-structural
//ff:what ST-1 positive — states/*.md 파싱 실패 시 [ST-1] 진단

package statemachine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
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
