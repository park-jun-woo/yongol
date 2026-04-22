//ff:func feature=cli type=reporter control=sequence
//ff:what formatStepLine — StepResult를 "✓ name  summary  (E errors, W warnings)" 형식으로 포매팅
package main

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

func formatStepLine(s validate.StepResult, errs, warns int) string {
	mark := statusMark(s.Status)
	parts := []string{fmt.Sprintf("%s %-20s", mark, s.Name)}
	if s.Summary != "" {
		parts = append(parts, s.Summary)
	}
	if errs > 0 || warns > 0 {
		parts = append(parts, fmt.Sprintf("(%d errors, %d warnings)", errs, warns))
	}
	return strings.Join(parts, "  ")
}
