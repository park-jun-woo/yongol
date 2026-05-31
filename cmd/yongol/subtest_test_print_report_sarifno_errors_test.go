//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestPrintReportSARIFNoErrors — NoErrors 서브테스트
package main

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

func subtestTestPrintReportSARIFNoErrors(t *testing.T) {

	r := &validate.Report{Steps: []validate.StepResult{
		{Name: "step", Status: validate.StatusPass},
	}}
	var buf bytes.Buffer
	errs, warns, err := printReportSARIF(&buf, r, "")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if errs != 0 || warns != 0 {
		t.Errorf("expected (0,0), got (%d,%d)", errs, warns)
	}

}
