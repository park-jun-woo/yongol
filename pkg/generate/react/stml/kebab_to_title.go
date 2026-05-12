//ff:func feature=stml-gen type=util control=iteration dimension=1 topic=string-convert
//ff:what 케밥케이스 페이지명에서 -page 접미사를 제거하고 Title Case로 변환한다
package stml

import "strings"

// kebabToTitle converts a kebab-case page name to a Title Case label.
// It strips the "-page" suffix if present, then title-cases each word.
// e.g. "workflows" → "Workflows", "workflow-detail" → "Workflow Detail",
// "my-reservations-page" → "My Reservations".
func kebabToTitle(name string) string {
	name = strings.TrimSuffix(name, "-page")
	parts := strings.Split(name, "-")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}
