//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseLayoutIfPresent — STML 미탐지 / layouts 디렉토리 부재 / 정상 파싱 분기 검증

package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseLayoutIfPresent_NoSTML(t *testing.T) {
	fs := &Fullstack{}
	parseLayoutIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.Layouts != nil {
		t.Fatalf("expected no Layouts when STML absent")
	}
}

func TestParseLayoutIfPresent_NoLayoutsDir(t *testing.T) {
	frontend := t.TempDir() // no layouts/ subdir
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindSTML: {Kind: KindSTML, Path: frontend, Presence: SSOTPopulated},
	}
	parseLayoutIfPresent(fs, has)
	if fs.Layouts != nil {
		t.Fatalf("expected no Layouts when layouts/ dir missing")
	}
}

func TestParseLayoutIfPresent_Present(t *testing.T) {
	frontend := t.TempDir()
	layoutDir := filepath.Join(frontend, "layouts")
	if err := os.MkdirAll(layoutDir, 0o755); err != nil {
		t.Fatal(err)
	}
	html := "<div>\n  <slot data-outlet />\n</div>\n"
	if err := os.WriteFile(filepath.Join(layoutDir, "app.html"), []byte(html), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindSTML: {Kind: KindSTML, Path: frontend, Presence: SSOTPopulated},
	}
	parseLayoutIfPresent(fs, has)
	if len(fs.Layouts) != 1 {
		t.Fatalf("expected 1 layout, got %d (diags=%+v)", len(fs.Layouts), fs.ParseDiagnostics)
	}
}
