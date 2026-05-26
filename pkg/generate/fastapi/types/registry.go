//ff:type feature=gen-fastapi type=adapter
//ff:what FastAPITypeRegistry — ir.TypeRegistry 의 FastAPI 구현 (PGFamily + BindOpts → Python/SQLAlchemy TypeBinding)

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// FastAPITypeRegistry implements ir.TypeRegistry for the FastAPI (Python
// + SQLAlchemy + Pydantic) backend. Each PGFamily maps to a SQLAlchemy
// column type (DBType), a Python/Pydantic API type (APIType), and the
// conversion expressions between DB and API layers.
type FastAPITypeRegistry struct{}

// Bind maps a PGFamily + BindOpts to an ir.TypeBinding containing
// SQLAlchemy column type, Python/Pydantic API type, and Py↔DB conversion
// expressions.
func (r *FastAPITypeRegistry) Bind(family typemap.PGFamily, opts ir.BindOpts) ir.TypeBinding {
	return bindFastAPI(family, opts)
}
