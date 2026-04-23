//ff:func feature=migration type=parser control=selection
//ff:what applyTypeParams — VARCHAR/CHAR 길이·NUMERIC(p,s) 를 CanonicalType 에 반영
package migration

import "strings"

// applyTypeParams fills in CanonicalType.Length / Precision / Scale from
// the raw parameter string inside parens.
func applyTypeParams(ct *CanonicalType, params string) {
	if params == "" {
		return
	}
	switch ct.Base {
	case "VARCHAR", "CHAR":
		ct.Length = parseIntSafe(strings.TrimSpace(params))
	case "NUMERIC":
		parts := strings.Split(params, ",")
		ct.Precision = parseIntSafe(strings.TrimSpace(parts[0]))
		if len(parts) > 1 {
			ct.Scale = parseIntSafe(strings.TrimSpace(parts[1]))
		}
	}
}
