//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindDate — FamilyDate → SQLAlchemy Date / Python date 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindDate produces the TypeBinding for DATE columns. SQLAlchemy: Date,
// Python: date. DB insert uses date.fromisoformat(), API/response
// conversions call .isoformat().
func bindDate(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyDate,
		NotNull:        opts.NotNull,
		DBType:         "Date",
		APIType:        nullableAPIType("date", opts.NotNull),
		APIImports:     []string{"from datetime import date"},
		ToDBExpr:       "date.fromisoformat({var})",
		ToAPIExpr:      "{row}.{field}.isoformat()",
		ToResponseExpr: "{var}.isoformat()",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
