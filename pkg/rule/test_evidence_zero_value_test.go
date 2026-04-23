//ff:func feature=rule type=test control=sequence
//ff:what TestEvidence_ZeroValue — 제로값 Evidence 의 모든 필드가 empty string 인지 확인

package rule

import "testing"

func TestEvidence_ZeroValue(t *testing.T) {
	var e Evidence
	if e.Rule != "" || e.Level != "" || e.Ref != "" || e.Message != "" {
		t.Fatalf("zero Evidence not empty: %+v", e)
	}
}
