//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParsePolicyIfPresent — Policy 미탐지(return) + 탐지 시 ParsedPolicies 설정
package yongol

import (
	"testing"
)

func TestParsePolicyIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parsePolicyIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.ParsedPolicies != nil {
		t.Fatalf("expected no ParsedPolicies when absent")
	}
}
