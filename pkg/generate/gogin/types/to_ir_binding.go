//ff:func feature=gen-gogin type=util control=sequence
//ff:what toIRBinding — GoTypeBinding → ir.TypeBinding 변환 (Go 전용 필드 → 백엔드 공통 필드)

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// toIRBinding converts a Go-specific GoTypeBinding to the backend-agnostic
// ir.TypeBinding. Go imports are shared across DB and API layers, so all
// imports are placed into APIImports (DB layer in Go uses the same import
// set).
func toIRBinding(family typemap.PGFamily, opts ir.BindOpts, gb GoTypeBinding) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         family,
		NotNull:        opts.NotNull,
		DBType:         gb.SqlcGoType,
		DBImports:      nil,
		APIType:        gb.ApiField,
		APIImports:     gb.Imports,
		ToDBExpr:       gb.InsertExpr,
		ToAPIExpr:      gb.ConvertExpr,
		ToResponseExpr: gb.ResponseExpr,
		NilCheckExpr:   gb.NilCheckExpr,
		Supported:      gb.Supported,
		DefaultLiteral: gb.DefaultLiteral,
	}
}
