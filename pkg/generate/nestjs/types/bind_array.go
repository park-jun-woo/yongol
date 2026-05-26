//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindArray — FamilyArray → Prisma T[] / TypeScript T[] 바인딩 (element head 기반)

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindArray produces the TypeBinding for PG array columns (TEXT[],
// BIGINT[], etc.). The element type is resolved from opts.ElementHead
// using arrayElementTSType. Only the four native families (Integer /
// Float / String / Boolean) are supported as array elements.
func bindArray(opts ir.BindOpts) ir.TypeBinding {
	elem, ok := arrayElementTSType(opts.ElementHead)
	if !ok {
		return ir.TypeBinding{
			Family:    typemap.FamilyArray,
			NotNull:   opts.NotNull,
			DBType:    "/* unsupported array element: " + opts.ElementHead + " */",
			APIType:   "/* unsupported array element: " + opts.ElementHead + " */",
			Supported: false,
		}
	}
	return ir.TypeBinding{
		Family:         typemap.FamilyArray,
		NotNull:        opts.NotNull,
		DBType:         elem.prisma + "[]",
		APIType:        nullableAPIType(elem.ts+"[]", opts.NotNull),
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
