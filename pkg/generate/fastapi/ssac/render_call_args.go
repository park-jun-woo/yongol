//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderCallArgs — FieldArg 배열 → 쉼표 구분 인자 문자열 생성

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderCallArgs produces a comma-separated argument list for a function call.
func renderCallArgs(args []ir.FieldArg) string {
	parts := make([]string, 0, len(args))
	for _, a := range args {
		parts = append(parts, renderArgValue(a))
	}
	return strings.Join(parts, ", ")
}
