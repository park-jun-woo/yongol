//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindNumeric — FamilyNumeric → SQLAlchemy Numeric / Python Decimal 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindNumeric produces the TypeBinding for numeric/decimal columns.
// SQLAlchemy: Numeric(precision, scale), Python: Decimal. DB insert uses
// Decimal() constructor, API/response conversions call str() for precision
// preservation.
func bindNumeric(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyNumeric,
		NotNull:        opts.NotNull,
		DBType:         "Numeric",
		DBImports:      []string{"from sqlalchemy import Numeric"},
		APIType:        nullableAPIType("Decimal", opts.NotNull),
		APIImports:     []string{"from decimal import Decimal"},
		ToDBExpr:       "Decimal({var})",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "str({var})",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
