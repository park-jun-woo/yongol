//ff:func feature=gen-fastapi type=util control=sequence
//ff:what mapPGToPython — PostgreSQL 타입 → Python 타입 어노테이션 변환

package models

import "strings"

// mapPGToPython maps a PostgreSQL raw type to a Python type annotation.
func mapPGToPython(rawType string, notNull bool) string {
	upper := strings.ToUpper(rawType)
	isArray := strings.HasSuffix(upper, "[]")
	if isArray {
		upper = strings.TrimSuffix(upper, "[]")
	}
	if idx := strings.Index(upper, "("); idx != -1 {
		upper = upper[:idx]
	}

	base := pyFamily(upper)
	if isArray {
		base = "list[" + base + "]"
	}
	if !notNull {
		return base + " | None"
	}
	return base
}
