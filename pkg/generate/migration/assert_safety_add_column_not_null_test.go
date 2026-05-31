//ff:func feature=migration type=test-helper control=selection
//ff:what assertSafetyAddColumnNotNull — safetyAddColumnNotNull 의 MIG-002 발생 여부 검증 헬퍼
package migration

import "testing"

// assertSafetyAddColumnNotNull asserts whether safetyAddColumnNotNull emits a
// single MIG-002 error for op.
func assertSafetyAddColumnNotNull(t *testing.T, op AddColumn, want bool) {
	t.Helper()
	issues := safetyAddColumnNotNull(op)
	switch {
	case want:
		if len(issues) != 1 || issues[0].RuleID != "MIG-002" || issues[0].Level != SafetyError {
			t.Errorf("got %+v, want one MIG-002 error", issues)
		}
	case issues != nil:
		t.Errorf("want nil, got %+v", issues)
	}
}
