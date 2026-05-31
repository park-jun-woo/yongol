//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseSSaCIfPresent — SSaC 미탐지(return) + 탐지 시 ServiceFuncs 설정
package yongol

import (
	"testing"
)

func TestParseSSaCIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseSSaCIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.ServiceFuncs != nil {
		t.Fatalf("expected no ServiceFuncs when absent")
	}
}
