//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseFeaturesIfPresent — features 미탐지(return) + 탐지 시 Features/FeatureTables 설정
package yongol

import (
	"testing"
)

func TestParseFeaturesIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseFeaturesIfPresent(fs, t.TempDir(), map[SSOTKind]DetectedSSOT{})
	if fs.Features != nil {
		t.Fatalf("expected no Features when absent, got %+v", fs.Features)
	}
}
