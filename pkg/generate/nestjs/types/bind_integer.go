//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindInteger — FamilyInteger → Prisma Int / TypeScript number 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindInteger produces the TypeBinding for integer columns (BIGINT, INT,
// SMALLINT, SERIAL). Prisma: Int, TypeScript: number. Values pass through
// without conversion.
func bindInteger(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyInteger,
		NotNull:        opts.NotNull,
		DBType:         "Int",
		APIType:        nullableAPIType("number", opts.NotNull),
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
