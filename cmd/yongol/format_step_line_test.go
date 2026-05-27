//ff:func feature=cli type=test control=sequence
//ff:what formatStepLine test — step 출력 포맷 검증

package main

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestFormatStepLine(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		s := validate.StepResult{Name: "ddl", Status: validate.StatusPass, Summary: "ok"}
		got := formatStepLine(s, 0, 0)
		if !strings.Contains(got, "ddl") {
			t.Errorf("expected 'ddl' in output, got '%s'", got)
		}
		if strings.Contains(got, "errors") {
			t.Errorf("should not contain errors when 0/0, got '%s'", got)
		}
	})
	t.Run("WithErrors", func(t *testing.T) {
		s := validate.StepResult{Name: "openapi", Status: validate.StatusFail}
		got := formatStepLine(s, 2, 1)
		if !strings.Contains(got, "(2 errors, 1 warnings)") {
			t.Errorf("expected error counts, got '%s'", got)
		}
	})
}
