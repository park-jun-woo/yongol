//ff:func feature=validate type=test-helper control=iteration dimension=2 topic=query-structural
//ff:what collectFiredQRules — diagnostics 에서 지정된 Q-* rule ID 중 발화된 집합 반환

package query

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// collectFiredQRules scans diags and returns a set[ruleID] → bool limited to
// the given ruleIDs, by substring match of `[ruleID]` on each Message.
func collectFiredQRules(diags []diagnostic.Diagnostic, ruleIDs []string) map[string]bool {
	fired := make(map[string]bool)
	for _, d := range diags {
		for _, id := range ruleIDs {
			if strings.Contains(d.Message, "["+id+"]") {
				fired[id] = true
			}
		}
	}
	return fired
}
