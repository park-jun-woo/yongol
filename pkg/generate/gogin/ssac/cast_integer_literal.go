//ff:func feature=gen-gogin type=util control=sequence
//ff:what castIntegerLiteral — 정수 리터럴에 Go 타입 캐스트 추가 (e.g. "1" → "int64(1)")

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/gogin/types"

// castIntegerLiteral adds a Go type cast to an integer literal so it
// matches the pgtypex bridge function's expected parameter type.
func (g *methodGen) castIntegerLiteral(rendered string, binding types.GoTypeBinding) string {
	if !isIntegerLiteralStr(rendered) {
		return rendered
	}
	goType := goIntTypeFromSqlcType(binding.SqlcGoType)
	if goType == "" {
		return rendered
	}
	return goType + "(" + rendered + ")"
}
