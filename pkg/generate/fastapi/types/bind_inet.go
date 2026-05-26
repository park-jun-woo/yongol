//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindInet — FamilyInet → SQLAlchemy INET / Python str 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindInet produces the TypeBinding for INET/CIDR columns. SQLAlchemy
// uses the PG-specific INET dialect type. Python: str. Values pass
// through without conversion.
func bindInet(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyInet,
		NotNull:        opts.NotNull,
		DBType:         "INET",
		DBImports:      []string{"from sqlalchemy.dialects.postgresql import INET"},
		APIType:        nullableAPIType("str", opts.NotNull),
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
