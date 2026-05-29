//ff:func feature=rule type=test-helper control=iteration dimension=1
//ff:what required 슬라이스가 email+age 를 포함하는지 assert — populate_constraint_fields 회귀 테스트 헬퍼

package ground

import (
	"testing"
)

// assertConstraintRequiredContainsEmailAndAge reports a test error when the
// required slice does not contain both "email" and "age".
func assertConstraintRequiredContainsEmailAndAge(t *testing.T, req []string) {
	t.Helper()
	saw := map[string]bool{}
	for _, n := range req {
		saw[n] = true
	}
	if !saw["email"] || !saw["age"] {
		t.Errorf("required missing entry: %v", req)
	}
}
