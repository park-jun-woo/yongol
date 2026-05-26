//ff:func feature=gen-nestjs type=util control=sequence
//ff:what pgToPrismaType — PostgreSQL 타입 → Prisma 스칼라 타입 변환

package prisma

import "strings"

// pgToPrismaType maps a PostgreSQL raw type token to a Prisma scalar type.
func pgToPrismaType(rawType string) string {
	upper := strings.ToUpper(rawType)
	isArray := strings.HasSuffix(upper, "[]")
	if isArray {
		upper = strings.TrimSuffix(upper, "[]")
	}
	if idx := strings.Index(upper, "("); idx != -1 {
		upper = upper[:idx]
	}
	base := mapPGFamily(upper)
	if isArray {
		return base + "[]"
	}
	return base
}
