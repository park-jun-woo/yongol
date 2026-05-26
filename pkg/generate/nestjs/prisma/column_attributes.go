//ff:func feature=gen-nestjs type=util control=selection
//ff:what columnAttributes — DDL Column → Prisma 속성 어노테이션 문자열 생성

package prisma

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// columnAttributes produces Prisma attribute annotations for a column.
func columnAttributes(col ddl.Column, colName string, primaryKey []string) string {
	var attrs []string

	if isPrimaryKey(colName, primaryKey) {
		attrs = append(attrs, "@id")
	}

	upper := strings.ToUpper(col.RawType)
	switch {
	case strings.HasPrefix(upper, "UUID") && col.HasDefault:
		attrs = append(attrs, "@default(uuid())")
	case strings.HasPrefix(upper, "TIMESTAMP") && col.HasDefault:
		attrs = append(attrs, "@default(now())")
	case strings.HasPrefix(upper, "SERIAL") || strings.HasPrefix(upper, "BIGSERIAL"):
		attrs = append(attrs, "@default(autoincrement())")
	case col.HasDefault:
		attrs = append(attrs, "@default("+prismaDefault(col)+")")
	}

	return strings.Join(attrs, " ")
}
