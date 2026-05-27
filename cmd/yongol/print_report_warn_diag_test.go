//ff:func feature=cli type=test control=sequence
//ff:what TestPrintReportWarnDiag — printReportWarnDiag 헬퍼 함수 검증

package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestPrintReportWarnDiag(t *testing.T) {
	d := printReportWarnDiag("[W-01] some warning")
	if d.Level != diagnostic.LevelWarning {
		t.Errorf("Level = %v, want LevelWarning", d.Level)
	}
	if d.File != "x.ssac" {
		t.Errorf("File = %q, want %q", d.File, "x.ssac")
	}
	if d.Line != 1 {
		t.Errorf("Line = %d, want 1", d.Line)
	}
	if d.Phase != diagnostic.PhaseValidate {
		t.Errorf("Phase = %v, want PhaseValidate", d.Phase)
	}
	if d.Message != "[W-01] some warning" {
		t.Errorf("Message = %q, want %q", d.Message, "[W-01] some warning")
	}
}
