//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindInteger — FamilyInteger → SQLAlchemy Integer / Python int 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindInteger produces the TypeBinding for integer columns (BIGINT, INT,
// SMALLINT, SERIAL). SQLAlchemy: Integer, Python: int. Values pass through
// without conversion.
func bindInteger(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyInteger,
		NotNull:        opts.NotNull,
		DBType:         "Integer",
		APIType:        nullableAPIType("int", opts.NotNull),
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
