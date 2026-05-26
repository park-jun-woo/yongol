//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindInterval — FamilyInterval → SQLAlchemy Interval / Python timedelta 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindInterval produces the TypeBinding for INTERVAL columns. SQLAlchemy:
// Interval, Python: timedelta. Values pass through without conversion.
func bindInterval(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyInterval,
		NotNull:        opts.NotNull,
		DBType:         "Interval",
		APIType:        nullableAPIType("timedelta", opts.NotNull),
		APIImports:     []string{"from datetime import timedelta"},
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
