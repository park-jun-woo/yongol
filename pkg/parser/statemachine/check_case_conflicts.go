//ff:func feature=statemachine type=parser control=iteration dimension=1
//ff:what 상태명 대소문자 충돌을 검사하여 Diagnostic 반환
package statemachine

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// checkCaseConflicts checks for case-insensitive state name conflicts.
func checkCaseConflicts(id, file string, stateSet map[string]bool, lines []string, mermaidStartLine int) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	lowerMap := make(map[string]string) // lowercase → first seen form
	for s := range stateSet {
		low := strings.ToLower(s)
		prev, exists := lowerMap[low]
		if exists && prev != s {
			conflictLine := findStateLine(lines, s, mermaidStartLine)
			diags = append(diags, diagnostic.Diagnostic{
				File:    file,
				Line:    conflictLine,
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("state name conflict in %s: %q and %q differ only in case", id, prev, s),
			})
		}
		lowerMap[low] = s
	}
	return diags
}
