//ff:func feature=gen-nestjs type=util control=sequence
//ff:what fieldName — FieldArg.Field 접근자에서 선행 점 제거한 필드명 추출

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// fieldName extracts the field accessor name from a FieldArg, stripping
// the leading dot.
func fieldName(a ir.FieldArg) string {
	if len(a.Field) > 0 && a.Field[0] == '.' {
		return a.Field[1:]
	}
	return a.Field
}
