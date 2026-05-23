//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what diagMessages — 진단 목록에서 메시지 문자열 추출

package agent

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// diagMessages extracts message strings from diagnostics.
func diagMessages(diags []diagnostic.Diagnostic) []string {
	msgs := make([]string, len(diags))
	for i, d := range diags {
		msgs[i] = d.Message
	}
	return msgs
}
