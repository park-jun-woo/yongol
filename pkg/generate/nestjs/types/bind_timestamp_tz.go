//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindTimestampTZ — FamilyTimestampTZ → Prisma DateTime / TypeScript string (ISO 8601) 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindTimestampTZ produces the TypeBinding for TIMESTAMPTZ columns. Prisma:
// DateTime, TypeScript: string (ISO 8601). DB insert wraps with new Date(),
// API/response conversions call .toISOString().
func bindTimestampTZ(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyTimestampTZ,
		NotNull:        opts.NotNull,
		DBType:         "DateTime",
		APIType:        nullableAPIType("string", opts.NotNull),
		ToDBExpr:       "new Date({var})",
		ToAPIExpr:      "{row}.{field}.toISOString()",
		ToResponseExpr: "{var}.toISOString()",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
