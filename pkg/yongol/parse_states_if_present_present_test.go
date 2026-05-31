//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseStatesIfPresent — States 미탐지(return) + 탐지 시 StateDiagrams 설정
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseStatesIfPresent_Present(t *testing.T) {
	dir := t.TempDir()
	body := "# Gig\n\n```mermaid\nstateDiagram-v2\n    [*] --> draft\n    draft --> open: PublishGig\n    open --> closed: CloseGig\n```\n"
	if err := os.WriteFile(filepath.Join(dir, "gig.md"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindStates: {Kind: KindStates, Path: dir, Presence: SSOTPopulated},
	}
	parseStatesIfPresent(fs, has)
	if len(fs.StateDiagrams) == 0 {
		t.Fatalf("expected StateDiagrams populated, diags=%+v", fs.ParseDiagnostics)
	}
}
