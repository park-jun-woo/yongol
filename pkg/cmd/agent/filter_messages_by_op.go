//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what filterMessagesByOp — 특정 operationId를 언급하는 메시지 필터링

package agent

import "strings"

// filterMessagesByOp returns diagnostic messages that mention the given operationId.
func filterMessagesByOp(messages []string, opID string) []string {
	var filtered []string
	for _, m := range messages {
		if strings.Contains(m, opID) {
			filtered = append(filtered, m)
		}
	}
	return filtered
}
