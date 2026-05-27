//ff:func feature=gen-nestjs type=util control=selection
//ff:what singularize — 영어 복수형 snake_case 테이블명 → 단수형 변환

package prisma

import "strings"

// singularize removes the English plural suffix from a snake_case table name.
func singularize(name string) string {
	switch {
	case strings.HasSuffix(name, "ies"):
		return name[:len(name)-3] + "y"
	case strings.HasSuffix(name, "sses"):
		return name[:len(name)-2]
	case strings.HasSuffix(name, "xes"):
		return name[:len(name)-2]
	case strings.HasSuffix(name, "s") && !strings.HasSuffix(name, "ss"):
		return name[:len(name)-1]
	default:
		return name
	}
}
