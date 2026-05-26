//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindTimestamp — FamilyTimestamp → SQLAlchemy DateTime(timezone=False) / Python datetime 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindTimestamp produces the TypeBinding for TIMESTAMP (without time zone)
// columns. SQLAlchemy: DateTime(timezone=False), Python: datetime. DB
// insert uses datetime.fromisoformat(), API/response conversions call
// .isoformat().
func bindTimestamp(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyTimestamp,
		NotNull:        opts.NotNull,
		DBType:         "DateTime(timezone=False)",
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
