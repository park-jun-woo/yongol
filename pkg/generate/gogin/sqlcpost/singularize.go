//ff:func feature=gen-gogin type=util control=sequence
//ff:what singularize — 간단 영어 단수형 변환 (ies→y, es 탈락, s 탈락)

package sqlcpost

import "strings"

// singularize is a minimal English singularizer sufficient for the SQL
// conventions we observe: trailing "ies" → "y", trailing "es" on sibilants
// → drop "es", trailing "s" → drop "s". Unknown cases return input.
func singularize(s string) string {
	if strings.HasSuffix(s, "ies") {
		return s[:len(s)-3] + "y"
	}
	if strings.HasSuffix(s, "ses") || strings.HasSuffix(s, "xes") ||
		strings.HasSuffix(s, "zes") || strings.HasSuffix(s, "ches") ||
		strings.HasSuffix(s, "shes") {
		return s[:len(s)-2]
	}
	if strings.HasSuffix(s, "s") {
		return s[:len(s)-1]
	}
	return s
}
