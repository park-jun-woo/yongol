//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseLayoutIfPresent — STML 미탐지 / layouts 디렉토리 부재 / 정상 파싱 분기 검증
package yongol

import (
	"testing"
)

func TestParseLayoutIfPresent_NoSTML(t *testing.T) {
	fs := &Fullstack{}
	parseLayoutIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.Layouts != nil {
		t.Fatalf("expected no Layouts when STML absent")
	}
}
