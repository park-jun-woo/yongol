//ff:func feature=gen-ir type=util control=selection
//ff:what ddlTableSingularIR -- 복수형 lower-snake 테이블명 → 단수형 (gogin 미러)

package ir

import "strings"

// ddlTableSingularIR desingularises a lower-snake table name. Mirrors
// gogin/ssac/ddlTableSingular for consistency.
func ddlTableSingularIR(name string) string {
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
