//ff:func feature=manifest type=test-helper control=iteration dimension=1
//ff:what rolesSentinelAnnotated — sentinel 스캔 결과에서 roles INSERT 의 Annotated/StartLine 검증 헬퍼
package ddl

import "testing"

// rolesSentinelAnnotated reports whether the scan results contain the roles
// table sentinel, asserting it is Annotated with a positive StartLine.
func rolesSentinelAnnotated(t *testing.T, got []SentinelScanResult) bool {
	t.Helper()
	for _, r := range got {
		if r.Table != "roles" {
			continue
		}
		if !r.Annotated {
			t.Errorf("roles INSERT should be Annotated")
		}
		if r.StartLine <= 0 {
			t.Errorf("StartLine = %d, want > 0", r.StartLine)
		}
		return true
	}
	return false
}
