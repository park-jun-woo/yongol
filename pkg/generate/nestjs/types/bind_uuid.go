//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindUUID — FamilyUUID → Prisma String @db.Uuid / TypeScript string 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindUUID produces the TypeBinding for UUID columns. Prisma: String
// @db.Uuid, TypeScript: string. Values pass through without conversion.
func bindUUID(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyUUID,
		NotNull:        opts.NotNull,
		DBType:         "String @db.Uuid",
		APIType:        nullableAPIType("string", opts.NotNull),
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
