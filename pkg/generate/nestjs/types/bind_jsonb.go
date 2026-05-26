//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindJSONB — FamilyJSONB → Prisma Json / TypeScript Record<string, unknown> 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindJSONB produces the TypeBinding for JSONB/JSON columns. Prisma: Json,
// TypeScript: Record<string, unknown>. Values pass through without
// conversion — Prisma handles JSON serialisation natively.
func bindJSONB(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyJSONB,
		NotNull:        opts.NotNull,
		DBType:         "Json",
		APIType:        nullableAPIType("Record<string, unknown>", opts.NotNull),
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
