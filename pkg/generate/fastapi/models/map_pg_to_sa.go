//ff:func feature=gen-fastapi type=util control=sequence
//ff:what mapPGToSA — PostgreSQL 타입 → SQLAlchemy 컬럼 타입 변환

package models

import "strings"

// mapPGToSA maps a PostgreSQL raw type to a SQLAlchemy column type.
func mapPGToSA(rawType string) string {
	upper := strings.ToUpper(rawType)
	isArray := strings.HasSuffix(upper, "[]")
	if isArray {
		upper = strings.TrimSuffix(upper, "[]")
	}
	if idx := strings.Index(upper, "("); idx != -1 {
		upper = upper[:idx]
	}

	base := mapSAFamily(upper)
	if isArray {
		return "ARRAY(" + base + ")"
	}
	return base
}
