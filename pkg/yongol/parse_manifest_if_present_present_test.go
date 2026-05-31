//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseManifestIfPresent — manifest 미탐지(return) + 탐지 시 Manifest 설정
package yongol

import (
	"path/filepath"
	"testing"
)

func TestParseManifestIfPresent_Present(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "manifest.yaml"), minimalManifest)
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindConfig: {Kind: KindConfig, Path: filepath.Join(root, "manifest.yaml"), Presence: SSOTPopulated},
	}
	parseManifestIfPresent(fs, root, has)
	if fs.Manifest == nil {
		t.Fatalf("expected Manifest to be set, diags=%+v", fs.ParseDiagnostics)
	}
}
