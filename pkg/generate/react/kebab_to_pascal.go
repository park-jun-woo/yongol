//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what kebab-case 문자열을 PascalCase로 변환한다

package react

import "strings"

// kebabToPascal converts a kebab-case string to PascalCase.
// e.g. "workflow-detail" → "WorkflowDetail"
func kebabToPascal(s string) string {
	parts := strings.Split(s, "-")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return sanitizeComponentName(strings.Join(parts, ""))
}
