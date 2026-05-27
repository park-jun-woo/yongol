//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what formatColumnList — 컬럼 이름 배열 → 쉼표 구분 인용 문자열 생성

package models

import (
	"fmt"
	"strings"
)

// formatColumnList produces a comma-separated list of quoted column names.
func formatColumnList(cols []string) string {
	parts := make([]string, len(cols))
	for i, c := range cols {
		parts[i] = fmt.Sprintf("\"%s\"", c)
	}
	return strings.Join(parts, ", ")
}
