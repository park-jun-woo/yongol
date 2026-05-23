//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what filterDiagsByOp — 특정 operationId를 포함하는 진단 필터링

package agent

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func filterDiagsByOp(diags []diagnostic.Diagnostic, opID string) []diagnostic.Diagnostic {
	var filtered []diagnostic.Diagnostic
	for _, d := range diags {
		if strings.Contains(d.Message, opID) {
			filtered = append(filtered, d)
		}
	}
	if len(filtered) == 0 {
		return diags
	}
	return filtered
}
