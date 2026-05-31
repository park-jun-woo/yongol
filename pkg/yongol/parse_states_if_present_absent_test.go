//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseStatesIfPresent — States 미탐지(return) + 탐지 시 StateDiagrams 설정
package yongol

import (
	"testing"
)

func TestParseStatesIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseStatesIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.StateDiagrams != nil {
		t.Fatalf("expected no StateDiagrams when absent")
	}
}
