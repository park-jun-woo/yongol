//ff:func feature=gen-gogin type=util control=sequence
//ff:what BuildWhat — //ff:what <한 줄 설명> 문자열 조립

package ffannot

import "strings"

// BuildWhat renders a //ff:what annotation line.
// Newlines in desc are collapsed to a single space so the output stays on one line.
// Returns an empty string when desc is empty — callers decide whether to skip.
func BuildWhat(desc string) string {
	desc = strings.TrimSpace(desc)
	if desc == "" {
		return ""
	}
	// Collapse any embedded newlines — //ff:what is single-line.
	desc = strings.ReplaceAll(desc, "\r\n", " ")
	desc = strings.ReplaceAll(desc, "\n", " ")
	return "//ff:what " + desc
}
