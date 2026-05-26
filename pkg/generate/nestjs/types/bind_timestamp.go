//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindTimestamp — FamilyTimestamp → Prisma DateTime / TypeScript string (ISO 8601) 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindTimestamp produces the TypeBinding for TIMESTAMP (without time zone)
// columns. Prisma: DateTime, TypeScript: string (ISO 8601). DB insert wraps
// with new Date(), API/response conversions call .toISOString().
func bindTimestamp(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyTimestamp,
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
