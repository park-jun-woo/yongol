//ff:func feature=cli-init type=test control=sequence
//ff:what TestWriteProjectIDBoundary — i==0/camelCase/acronym/무경계에서 언더스코어 삽입 여부 검증
package cliinit

import (
	"strings"
)

func boundaryAt(s string, i int) string {
	runes := []rune(s)
	var b strings.Builder
	writeProjectIDBoundary(&b, runes, i, runes[i])
	return b.String()
}
