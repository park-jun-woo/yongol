//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindDate — FamilyDate → Prisma DateTime @db.Date / TypeScript string 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindDate produces the TypeBinding for DATE columns. Prisma: DateTime
// @db.Date, TypeScript: string. DB insert wraps with new Date(),
// API/response conversions call .toISOString().slice(0,10) to extract the
// date portion.
func bindDate(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyDate,
		NotNull:        opts.NotNull,
		DBType:         "DateTime @db.Date",
		APIType:        nullableAPIType("string", opts.NotNull),
		ToDBExpr:       "new Date({var})",
		ToAPIExpr:      "{row}.{field}.toISOString().slice(0,10)",
		ToResponseExpr: "{var}.toISOString().slice(0,10)",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
