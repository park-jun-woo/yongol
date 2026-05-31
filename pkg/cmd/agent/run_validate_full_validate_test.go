//ff:func feature=agent type=test control=sequence
//ff:what TestRunValidate — DetectSSOTs 에러 / 파싱 진단 / 정상 validate 진단 수집 분기 검증
package agent

import (
	"testing"
)

func TestRunValidateFullValidate(t *testing.T) {
	// An empty (but valid) specs directory parses without diagnostics and runs
	// the full validate pass, exercising the report-aggregation loop.
	dir := t.TempDir()
	if _, err := runValidate(dir); err != nil {
		t.Fatalf("unexpected error on empty specs: %v", err)
	}
}
