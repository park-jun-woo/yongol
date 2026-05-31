//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseScenarioIfPresent — Hurl 미탐지(return) + 탐지 시 HurlFiles/HurlEntries 설정
package yongol

import (
	"testing"
)

func TestParseScenarioIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseScenarioIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.HurlFiles != nil || fs.HurlEntries != nil {
		t.Fatalf("expected no Hurl data when absent")
	}
}
