//ff:func feature=cli type=reporter control=iteration dimension=1
//ff:what printLevelDiags — 특정 레벨 diagnostic 을 "[step]" 헤더와 함께 출력
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// printLevelDiags writes diagnostics of the given level under a "[step]" header.
// Multi-line messages and " → 권고: ..." suggestions are indented for readability.
func printLevelDiags(w io.Writer, s validate.StepResult, level diagnostic.Level) {
	first := true
	for _, d := range s.Diagnostics {
		if d.Level != level {
			continue
		}
		if first {
			fmt.Fprintf(w, "[%s]\n", s.Name)
			first = false
		}
		// Advice 필드를 우선 사용. 비어있으면 Message 안 인라인된 → 권고: 분리 (Phase003 마이그레이션 중인 함수 호환).
		main := d.Message
		advice := d.Advice
		if advice == "" {
			main, advice = splitAdvice(d.Message)
		}
		fmt.Fprintf(w, "  - %s%s\n", formatLocation(d.File, d.Line), main)
		if advice != "" {
			fmt.Fprintf(w, "      ↳ 권고: %s\n", advice)
		}
	}
}
