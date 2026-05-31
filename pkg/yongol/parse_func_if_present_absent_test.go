//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseFuncIfPresent — Func 미탐지(return) + 탐지 시 ProjectFuncSpecs 설정
package yongol

import (
	"testing"
)

func TestParseFuncIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseFuncIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.ProjectFuncSpecs != nil {
		t.Fatalf("expected no specs when Func absent, got %+v", fs.ProjectFuncSpecs)
	}
}
