//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what pascalCaseSnake — snake_case 문자열을 언더스코어 기준 PascalCase로 조합

package ssac

import "strings"

// pascalCaseSnake joins snake_case segments into PascalCase. Empty segments
// (caused by leading/trailing/duplicate underscores) are skipped. Caller
// (pascalCase) guarantees s contains at least one underscore.
func pascalCaseSnake(s string) string {
	parts := strings.Split(s, "_")
	var b strings.Builder
	for _, p := range parts {
		if p != "" {
			b.WriteString(strings.ToUpper(p[:1]) + p[1:])
		}
	}
	return b.String()
}
