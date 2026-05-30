//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseFeaturesIfPresent — features 미탐지(return) + 탐지 시 Features/FeatureTables 설정

package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFeaturesIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseFeaturesIfPresent(fs, t.TempDir(), map[SSOTKind]DetectedSSOT{})
	if fs.Features != nil {
		t.Fatalf("expected no Features when absent, got %+v", fs.Features)
	}
}

func TestParseFeaturesIfPresent_Present(t *testing.T) {
	root := t.TempDir()
	data := "features:\n  - op: CreateWorkflow\n    path: POST /workflows\n    desc: Create a new workflow\n"
	if err := os.WriteFile(filepath.Join(root, "features.yaml"), []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindFeatures: {Kind: KindFeatures, Path: filepath.Join(root, "features.yaml"), Presence: SSOTPopulated},
	}
	parseFeaturesIfPresent(fs, root, has)
	if len(fs.Features) != 1 {
		t.Fatalf("expected 1 feature, got %d", len(fs.Features))
	}
}
