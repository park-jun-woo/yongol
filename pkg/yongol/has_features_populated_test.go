//ff:func feature=orchestrator type=test-helper control=iteration dimension=1
//ff:what hasFeaturesPopulated — DetectSSOTs 결과에 SSOTPopulated KindFeatures 항목 존재 여부 검사 헬퍼
package yongol

import "testing"

// hasFeaturesPopulated reports whether found contains a KindFeatures entry,
// asserting it is SSOTPopulated when present.
func hasFeaturesPopulated(t *testing.T, found []DetectedSSOT) bool {
	t.Helper()
	for _, d := range found {
		if d.Kind != KindFeatures {
			continue
		}
		if d.Presence != SSOTPopulated {
			t.Errorf("features presence = %v, want SSOTPopulated", d.Presence)
		}
		return true
	}
	return false
}
