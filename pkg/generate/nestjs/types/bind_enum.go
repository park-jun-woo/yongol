//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindEnum — FamilyEnum → Prisma String / TypeScript string 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindEnum produces the TypeBinding for enum columns (CHECK IN constraint).
// Prisma: String, TypeScript: string. Values pass through without
// conversion. The enum values from DDL CHECK IN constraints are available
// via opts.EnumValues but are not embedded in the binding — downstream
// Prisma schema / DTO generators use them separately.
func bindEnum(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyEnum,
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
