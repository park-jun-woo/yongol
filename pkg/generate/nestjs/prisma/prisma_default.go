//ff:func feature=gen-nestjs type=util control=sequence
//ff:what prismaDefault — DDL DEFAULT 리터럴 → Prisma default 표현식 변환

package prisma

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// prismaDefault converts a DDL default literal to a Prisma default expression.
func prismaDefault(col ddl.Column) string {
	lit := col.DefaultLiteral
	if lit == "" {
		return "\"\""
	}
	upper := strings.ToUpper(lit)
	if upper == "NOW()" || upper == "CURRENT_TIMESTAMP" {
		return "now()"
	}
	if upper == "GEN_RANDOM_UUID()" || upper == "UUID_GENERATE_V4()" {
		return "uuid()"
	}
	if upper == "TRUE" || upper == "FALSE" {
		return strings.ToLower(lit)
	}
	if _, err := fmt.Sscanf(lit, "%f", new(float64)); err == nil {
		return lit
	}
	return fmt.Sprintf("%q", lit)
}
