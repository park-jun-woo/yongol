//ff:func feature=util type=util control=iteration dimension=1 topic=string-convert
//ff:what KebabToCamel — kebab-case → camelCase (dash 없는 문자열은 그대로)

package caseconv

import "strings"

// KebabToCamel converts kebab-case to camelCase. Strings without '-' are
// returned unchanged. Example: "data-fetch" → "dataFetch".
func KebabToCamel(s string) string {
	if !strings.Contains(s, "-") {
		return s
	}
	parts := strings.Split(s, "-")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}
