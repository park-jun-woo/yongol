//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindFloat — FamilyFloat → SQLAlchemy Float / Python float 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindFloat produces the TypeBinding for float columns (REAL, FLOAT4,
// FLOAT8). SQLAlchemy: Float, Python: float. Values pass through without
// conversion.
func bindFloat(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyFloat,
		NotNull:        opts.NotNull,
		DBType:         "Float",
		APIType:        nullableAPIType("float", opts.NotNull),
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
