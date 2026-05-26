//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindInet — FamilyInet → Prisma String / TypeScript string 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindInet produces the TypeBinding for INET/CIDR columns. Prisma has no
// native type for network addresses, so String is used. TypeScript: string.
// Values pass through without conversion.
func bindInet(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyInet,
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
