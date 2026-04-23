//ff:func feature=report type=util control=sequence topic=sarif
//ff:what extractRuleID — diagnostic 메시지 앞 "[RULE-ID]" 토큰을 (ruleID, remainingMessage) 로 분리
package sarif

import (
	"regexp"
	"strings"
)

// ruleIDPattern matches a leading "[RULE-ID]" prefix in diagnostic messages,
// e.g. "[S-27] Variable foo not declared" or "[XOS-15] ...".
// Rule IDs: uppercase letters + dash + digits (S-27, XOS-15, Q-10, M-2, …).
var ruleIDPattern = regexp.MustCompile(`^\[([A-Z]+(?:-[A-Z]+)*-\d+)\]\s*`)

// extractRuleID pulls a leading "[RULE-ID]" token out of a diagnostic message.
// Returns (ruleID, remainingMessage). When no prefix is present the original
// message is returned with an empty ruleID.
func extractRuleID(msg string) (string, string) {
	m := ruleIDPattern.FindStringSubmatchIndex(msg)
	if m == nil {
		return "", msg
	}
	id := msg[m[2]:m[3]]
	rest := strings.TrimSpace(msg[m[1]:])
	return id, rest
}
