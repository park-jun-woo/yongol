//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseSTMLIfPresent — STML 미탐지(return) + 탐지 시 STMLPages 설정

package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseSTMLIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseSTMLIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.STMLPages != nil {
		t.Fatalf("expected no STMLPages when absent")
	}
}

func TestParseSTMLIfPresent_Present(t *testing.T) {
	dir := t.TempDir()
	page := "<main>\n  <h1>Home</h1>\n</main>\n"
	if err := os.WriteFile(filepath.Join(dir, "home.html"), []byte(page), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindSTML: {Kind: KindSTML, Path: dir, Presence: SSOTPopulated},
	}
	parseSTMLIfPresent(fs, has)
	if len(fs.STMLPages) != 1 {
		t.Fatalf("expected 1 STML page, got %d (diags=%+v)", len(fs.STMLPages), fs.ParseDiagnostics)
	}
}
