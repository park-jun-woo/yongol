//ff:func feature=cli type=test control=sequence
//ff:what printReportErrDiag test — 헬퍼 빌더 결과 검증

package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestPrintReportErrDiag(t *testing.T) {
	d := printReportErrDiag("X-01: foo")
	if d.Level != diagnostic.LevelError {
		t.Errorf("expected LevelError, got %q", d.Level)
	}
	if d.Message != "X-01: foo" {
		t.Errorf("expected message, got %q", d.Message)
	}
	if d.File == "" {
		t.Error("expected non-empty file")
	}
}
