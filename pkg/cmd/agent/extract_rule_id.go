//ff:func feature=agent type=helper control=sequence
//ff:what extractRuleID — 진단 메시지에서 룰 ID 추출 ([S-74] 또는 S-74: 형식)

package agent

import "strings"

// extractRuleID extracts the rule ID from a diagnostic message.
// Handles both "[S-74] ..." and "S-74: ..." formats.
func extractRuleID(msg string) string {
	// Try bracket format first: [S-74]
	if strings.HasPrefix(msg, "[") {
		end := strings.Index(msg, "]")
		if end > 1 && end <= 20 {
			return msg[1:end]
		}
	}
	// Try colon format: S-74:
	idx := strings.Index(msg, ":")
	if idx > 0 && idx <= 20 {
		candidate := strings.TrimSpace(msg[:idx])
		if len(candidate) >= 2 && len(candidate) <= 15 {
			return candidate
		}
	}
	return ""
}
