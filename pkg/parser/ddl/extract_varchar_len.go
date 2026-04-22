//ff:func feature=manifest type=util control=sequence
//ff:what extractVarcharLen — VARCHAR(N) 타입에서 길이 N 추출
package ddl

import (
	"regexp"
	"strconv"
)

var reVarcharLen = regexp.MustCompile(`(?i)VARCHAR\((\d+)\)`)

func extractVarcharLen(colType string) int {
	m := reVarcharLen.FindStringSubmatch(colType)
	if len(m) == 2 {
		n, _ := strconv.Atoi(m[1])
		return n
	}
	return 0
}
