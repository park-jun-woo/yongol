//ff:func feature=chain type=test control=sequence
//ff:what buildStateLink 가 transition 유무에 따라 line/summary 를 구성하는지 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildStateLink(t *testing.T) {
	specsDir := t.TempDir()
	statesDir := filepath.Join(specsDir, "states")
	if err := os.MkdirAll(statesDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	content := "stateDiagram-v2\n  Pending --> Cancelled: cancel\n"
	if err := os.WriteFile(filepath.Join(statesDir, "reservation.md"), []byte(content), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	// With a transition that is found in the file.
	link := buildStateLink("reservation", specsDir, map[string]string{"reservation": "cancel"})
	if link.Kind != "StateDiag" || link.File != "states/reservation.md" {
		t.Errorf("link fields: %+v", link)
	}
	if link.Line != 2 {
		t.Errorf("line: got %d, want 2", link.Line)
	}
	if link.Summary != "diagram: reservation -> cancel" {
		t.Errorf("summary: %q", link.Summary)
	}

	// Without a transition: line stays 0, summary has no arrow.
	linkNoTrans := buildStateLink("reservation", specsDir, map[string]string{})
	if linkNoTrans.Line != 0 {
		t.Errorf("no-transition line: got %d, want 0", linkNoTrans.Line)
	}
	if linkNoTrans.Summary != "diagram: reservation" {
		t.Errorf("no-transition summary: %q", linkNoTrans.Summary)
	}
}
