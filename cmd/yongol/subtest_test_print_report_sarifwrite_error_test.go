//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestPrintReportSARIFWriteError — WriteError 서브테스트
package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

func subtestTestPrintReportSARIFWriteError(t *testing.T) {

	r := &validate.Report{Steps: []validate.StepResult{
		{Name: "step", Status: validate.StatusPass},
	}}
	_, _, err := printReportSARIF(&failWriter{}, r, "")
	if err == nil {
		t.Fatal("expected write error")
	}

}
