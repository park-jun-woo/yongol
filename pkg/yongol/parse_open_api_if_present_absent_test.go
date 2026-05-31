//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseOpenAPIIfPresent — 미탐지 / 로드 에러 / 정상 로드 분기 검증
package yongol

import (
	"testing"
)

func TestParseOpenAPIIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseOpenAPIIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.OpenAPIDoc != nil {
		t.Fatalf("expected no OpenAPIDoc when absent")
	}
}
