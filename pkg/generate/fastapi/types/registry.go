//ff:type feature=gen-fastapi type=adapter
//ff:what FastAPITypeRegistry — ir.TypeRegistry 의 FastAPI 스켈레톤 구현 (미구현: Supported=false)

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// FastAPITypeRegistry implements ir.TypeRegistry for the FastAPI (Python
// + SQLAlchemy) backend. This is a skeleton — all families return
// Supported=false until the FastAPI backend is fully implemented.
type FastAPITypeRegistry struct{}

// Bind always returns an unsupported TypeBinding. The FastAPI backend type
// mappings (Python type names, SQLAlchemy column types, Py↔DB conversion
// expressions) will be filled in when the FastAPI code generator is built.
func (r *FastAPITypeRegistry) Bind(family typemap.PGFamily, opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:    family,
		NotNull:   opts.NotNull,
		DBType:    "/* fastapi: not yet implemented */",
		APIType:   "/* fastapi: not yet implemented */",
		Supported: false,
	}
}
