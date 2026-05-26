//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderSAWhere — FieldArg 배열 → SQLAlchemy .where() 절 문자열 생성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderSAWhere builds a SQLAlchemy .where() clause from FieldArgs.
func renderSAWhere(model string, args []ir.FieldArg) string {
	if len(args) == 0 {
		return ""
	}
	pyModel := pascalCase(model)
	parts := make([]string, 0, len(args))
	for _, a := range args {
		key := resolveArgKey(a)
		parts = append(parts, fmt.Sprintf("%s.%s == %s", pyModel, key, renderArgValue(a)))
	}
	return ".where(" + strings.Join(parts, ", ") + ")"
}
