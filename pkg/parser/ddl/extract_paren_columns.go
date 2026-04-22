//ff:func feature=manifest type=util control=iteration dimension=1
//ff:what extractParenColumns — 괄호 안 컬럼 목록 추출
package ddl

import "strings"

func extractParenColumns(line string) []string {
	parenIdx := strings.Index(line, "(")
	if parenIdx < 0 {
		return nil
	}
	inner := line[parenIdx+1:]
	inner = strings.TrimSuffix(strings.TrimSpace(inner), ",")
	inner = strings.TrimSuffix(inner, ");")
	inner = strings.TrimSuffix(inner, ")")
	var cols []string
	for _, c := range strings.Split(inner, ",") {
		c = strings.TrimSpace(c)
		if c != "" {
			cols = append(cols, c)
		}
	}
	return cols
}
