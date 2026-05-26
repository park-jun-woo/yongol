//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindUUID — FamilyUUID → SQLAlchemy Uuid / Python uuid.UUID 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindUUID produces the TypeBinding for UUID columns. SQLAlchemy 2.0: Uuid,
// Python: uuid.UUID. Values pass through without conversion.
func bindUUID(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyUUID,
		NotNull:        opts.NotNull,
		DBType:         "Uuid",
		DBImports:      []string{"from sqlalchemy import Uuid"},
		APIType:        nullableAPIType("uuid.UUID", opts.NotNull),
		APIImports:     []string{"import uuid"},
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
