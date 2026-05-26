//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindNumeric — FamilyNumeric → Prisma Decimal / TypeScript string 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindNumeric produces the TypeBinding for numeric/decimal columns. Prisma:
// Decimal, TypeScript: string (precision preservation). DB insert wraps
// with new Decimal(), API/response conversions call .toString().
func bindNumeric(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyNumeric,
		NotNull:        opts.NotNull,
		DBType:         "Decimal",
		DBImports:      []string{"Decimal from @prisma/client/runtime/library"},
		APIType:        nullableAPIType("string", opts.NotNull),
		ToDBExpr:       "new Decimal({var})",
		ToAPIExpr:      "{row}.{field}.toString()",
		ToResponseExpr: "{var}.toString()",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
