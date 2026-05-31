//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseManifestIfPresent — manifest 미탐지(return) + 탐지 시 Manifest 설정
package yongol

import (
	"testing"
)

func TestParseManifestIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseManifestIfPresent(fs, t.TempDir(), map[SSOTKind]DetectedSSOT{})
	if fs.Manifest != nil {
		t.Fatalf("expected no Manifest when absent")
	}
}
