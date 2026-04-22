//ff:func feature=manifest type=util control=iteration dimension=1
//ff:what parseCheckEnum — CHECK (col IN (...)) 절에서 컬럼명과 허용 값 파싱
package ddl

import (
	"regexp"
	"strings"
)

var reCheckEnum = regexp.MustCompile(`(?i)CHECK\s*\(\s*(\w+)\s+IN\s*\(([^)]+)\)\s*\)`)

func parseCheckEnum(line string) (string, []string) {
	m := reCheckEnum.FindStringSubmatch(line)
	if len(m) < 3 {
		return "", nil
	}
	var vals []string
	for _, v := range strings.Split(m[2], ",") {
		v = strings.TrimSpace(v)
		v = strings.Trim(v, "'\"")
		if v != "" {
			vals = append(vals, v)
		}
	}
	return m[1], vals
}
