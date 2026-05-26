//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindBytea — FamilyBytea → SQLAlchemy LargeBinary / Python bytes 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindBytea produces the TypeBinding for BYTEA columns. SQLAlchemy:
// LargeBinary, Python: bytes. Values pass through without conversion.
func bindBytea(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyBytea,
		NotNull:        opts.NotNull,
		DBType:         "LargeBinary",
		APIType:        nullableAPIType("bytes", opts.NotNull),
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
