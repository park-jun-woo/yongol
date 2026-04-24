//ff:func feature=gen-gogin type=util control=selection
//ff:what ddlTableSingular — 복수형 lower-snake 테이블명 → 단수형

package ssac

import (
	"strings"
)

// ddlTableSingular desingularises a lower-snake table name to the sqlc
// model name lower form. Kept simple — matches inflection.Singular on the
// zenflow fixture (users / organizations / workflows / actions /
// execution_logs).
func ddlTableSingular(name string) string {
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
