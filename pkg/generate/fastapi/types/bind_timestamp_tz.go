//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindTimestampTZ — FamilyTimestampTZ → SQLAlchemy DateTime(timezone=True) / Python datetime 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindTimestampTZ produces the TypeBinding for TIMESTAMPTZ columns.
// SQLAlchemy: DateTime(timezone=True), Python: datetime. DB insert uses
// datetime.fromisoformat(), API/response conversions call .isoformat().
func bindTimestampTZ(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyTimestampTZ,
		NotNull:        opts.NotNull,
		DBType:         "DateTime(timezone=True)",
		APIType:        nullableAPIType("datetime", opts.NotNull),
		APIImports:     []string{"from datetime import datetime"},
		ToDBExpr:       "datetime.fromisoformat({var})",
		ToAPIExpr:      "{row}.{field}.isoformat()",
		ToResponseExpr: "{var}.isoformat()",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
