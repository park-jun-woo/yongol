//ff:func feature=gen-fastapi type=util control=sequence
//ff:what saDefault — DDL DEFAULT 리터럴 → SQLAlchemy default 표현식 변환

package models

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// saDefault converts a DDL default literal to a SQLAlchemy default expression.
func saDefault(col ddl.Column) string {
	lit := col.DefaultLiteral
	upper := strings.ToUpper(col.RawType)

	if strings.HasPrefix(upper, "UUID") && col.HasDefault {
		return "default=uuid.uuid4"
	}
	if (strings.HasPrefix(upper, "TIMESTAMP") || strings.HasPrefix(upper, "DATE")) && col.HasDefault {
		return "server_default=\"now()\""
	}
	if strings.HasPrefix(upper, "SERIAL") || strings.HasPrefix(upper, "BIGSERIAL") {
		return "" // auto-increment handled by SA Integer + primary_key
	}
	if lit == "" {
		return ""
	}
	upperLit := strings.ToUpper(lit)
	if upperLit == "TRUE" || upperLit == "FALSE" {
		lower := strings.ToLower(lit)
		titled := strings.ToUpper(lower[:1]) + lower[1:]
		return fmt.Sprintf("default=%s", titled)
	}
	if _, err := fmt.Sscanf(lit, "%f", new(float64)); err == nil {
		return fmt.Sprintf("default=%s", lit)
	}
	return fmt.Sprintf("default=\"%s\"", lit)
}
