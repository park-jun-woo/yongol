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
	var hasFeatures bool
	for _, d := range found {
		if d.Kind == KindFeatures {
			hasFeatures = true
			if d.Presence != SSOTPopulated {
				t.Errorf("features presence = %v, want SSOTPopulated", d.Presence)
			}
		}
	}
	if !hasFeatures {
		t.Fatalf("expected a KindFeatures entry, got %+v", found)
	}
}
