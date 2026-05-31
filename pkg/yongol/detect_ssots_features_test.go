//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — features.yaml 존재 시 KindFeatures 항목이 감지되는지 검증
package yongol

import (
	"path/filepath"
	"testing"
)

func TestDetectSSOTsFeaturesPresent(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "features.yaml"), "features: []\n")

	found, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs failed: %v", err)
	}
	if !hasFeaturesPopulated(t, found) {
		t.Fatalf("expected a KindFeatures entry, got %+v", found)
	}
}
