//ff:func feature=gen-fastapi type=util control=sequence
//ff:what bindJSONB — FamilyJSONB → SQLAlchemy JSONB / Python dict[str, Any] 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindJSONB produces the TypeBinding for JSONB/JSON columns. SQLAlchemy:
// JSONB (PG dialect), Python: dict[str, Any]. Values pass through without
// conversion — SQLAlchemy handles JSON serialisation natively.
func bindJSONB(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyJSONB,
		NotNull:        opts.NotNull,
		DBType:         "JSONB",
		DBImports:      []string{"from sqlalchemy.dialects.postgresql import JSONB"},
		APIType:        nullableAPIType("dict[str, Any]", opts.NotNull),
		APIImports:     []string{"from typing import Any"},
		ToDBExpr:       "{var}",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
