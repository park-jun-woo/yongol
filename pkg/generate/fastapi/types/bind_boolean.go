//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindBoolean — FamilyBoolean → SQLAlchemy Boolean / Python bool 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindBoolean produces the TypeBinding for boolean columns (BOOLEAN,
// BOOL). SQLAlchemy: Boolean, Python: bool. Values pass through without
// conversion.
func bindBoolean(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyBoolean,
		NotNull:        opts.NotNull,
		DBType:         "Boolean",
		APIType:        nullableAPIType("bool", opts.NotNull),
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
