//ff:func feature=cli type=test control=sequence
//ff:what printLevelDiags level filter test — 불일치 레벨 진단 필터링 검증

package main

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestPrintLevelDiagsFiltersMismatchedLevel(t *testing.T) {
	var buf bytes.Buffer
	step := validate.StepResult{
		Name: "filter-test",
		Diagnostics: []diagnostic.Diagnostic{
			{Level: diagnostic.LevelWarning, Message: "W-01: something"},
		},
	}
	printLevelDiags(&buf, step, diagnostic.LevelError)
	if buf.Len() != 0 {
		t.Errorf("expected no output when no matching level, got: %q", buf.String())
	}
}
