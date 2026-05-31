//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseLayoutIfPresent — STML 미탐지 / layouts 디렉토리 부재 / 정상 파싱 분기 검증
package yongol

import (
	"testing"
)

func TestParseLayoutIfPresent_NoLayoutsDir(t *testing.T) {
	frontend := t.TempDir() // no layouts/ subdir
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindSTML: {Kind: KindSTML, Path: frontend, Presence: SSOTPopulated},
	}
	parseLayoutIfPresent(fs, has)
	if fs.Layouts != nil {
		t.Fatalf("expected no Layouts when layouts/ dir missing")
	}
}
