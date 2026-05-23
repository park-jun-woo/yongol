//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what collectRemainingFiles — 남은 에러의 파일:줄 요약 목록 반환

package agent

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// collectRemainingFiles returns sorted unique file:line summaries for remaining errors.
func collectRemainingFiles(diags []diagnostic.Diagnostic, absSpecs string) []string {
	seen := map[string]bool{}
	var result []string
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			continue
		}
		rel := rebaseFile(d.File, absSpecs)
		entry := fmt.Sprintf("%s:%d %s", rel, d.Line, d.Message)
		if seen[entry] {
			continue
		}
		seen[entry] = true
		result = append(result, entry)
	}
	sort.Strings(result)
	return result
}
