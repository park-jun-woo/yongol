//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseScenarioIfPresent — Hurl 미탐지(return) + 탐지 시 HurlFiles/HurlEntries 설정

package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseScenarioIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseScenarioIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.HurlFiles != nil || fs.HurlEntries != nil {
		t.Fatalf("expected no Hurl data when absent")
	}
}

func TestParseScenarioIfPresent_Present(t *testing.T) {
	dir := t.TempDir()
	content := "GET http://localhost:8080/workflows\nHTTP 200\n"
	if err := os.WriteFile(filepath.Join(dir, "scenario.hurl"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindScenario: {Kind: KindScenario, Path: dir, Presence: SSOTPopulated},
	}
	parseScenarioIfPresent(fs, has)
	if len(fs.HurlFiles) == 0 {
		t.Fatalf("expected HurlFiles populated")
	}
	if len(fs.HurlEntries) == 0 {
		t.Fatalf("expected HurlEntries populated")
	}
}
