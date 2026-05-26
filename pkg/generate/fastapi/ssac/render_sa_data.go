//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderSAData — FieldArg 배열 → SQLAlchemy 모델 keyword argument 문자열 생성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderSAData builds SQLAlchemy model keyword arguments from FieldArgs.
func renderSAData(args []ir.FieldArg) string {
	if len(args) == 0 {
		return "**body"
	}
	parts := make([]string, 0, len(args))
	for _, a := range args {
		key := resolveDataKey(a)
		if key == "" {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s=%s", key, renderArgValue(a)))
	}
	if len(parts) == 0 {
		return "**body"
	}
	return strings.Join(parts, ", ")
}
