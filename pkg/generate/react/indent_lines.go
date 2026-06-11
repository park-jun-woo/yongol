//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what indentLines — 각 라인 앞에 들여쓰기를 붙여 개행으로 연결

package react

import "strings"

// indentLines prefixes each line with indent and joins them with trailing
// newlines — the claims store emitter renders the same state-object body
// at the persist (6-space) and memory (2-space) depths from one source.
func indentLines(lines []string, indent string) string {
	var sb strings.Builder
	for _, l := range lines {
		sb.WriteString(indent)
		sb.WriteString(l)
		sb.WriteString("\n")
	}
	return sb.String()
}
