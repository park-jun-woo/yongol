//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindInterval — FamilyInterval → Prisma String / TypeScript string 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindInterval produces the TypeBinding for INTERVAL columns. Prisma has
// no native type for intervals, so String is used. TypeScript: string.
// Values pass through without conversion.
func bindInterval(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyInterval,
		NotNull:        opts.NotNull,
		DBType:         "String",
		APIType:        nullableAPIType("string", opts.NotNull),
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
