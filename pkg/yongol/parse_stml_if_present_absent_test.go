//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseSTMLIfPresent — STML 미탐지(return) + 탐지 시 STMLPages 설정
package yongol

import (
	"testing"
)

func TestParseSTMLIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseSTMLIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.STMLPages != nil {
		t.Fatalf("expected no STMLPages when absent")
	}
}
