//ff:func feature=cli type=reporter control=sequence
//ff:what formatStepLine — formats a StepResult as "✓ name  summary  (E errors, W warnings)"
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
