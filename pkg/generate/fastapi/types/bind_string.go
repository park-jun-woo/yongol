//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindString — FamilyString → SQLAlchemy String / Python str 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindString produces the TypeBinding for string/text columns (VARCHAR,
// TEXT, CHAR, BPCHAR). SQLAlchemy: String, Python: str. Values pass
// through without conversion.
func bindString(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyString,
		NotNull:        opts.NotNull,
		DBType:         "String",
		APIType:        nullableAPIType("str", opts.NotNull),
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
