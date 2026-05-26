//ff:type feature=gen-gogin type=adapter
//ff:what GoTypeRegistry — ir.TypeRegistry 의 Go+Gin 구현 (PGFamily + BindOpts → ir.TypeBinding)

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// GoTypeRegistry implements ir.TypeRegistry for the Go+Gin backend. It
// wraps the existing family binding functions (nativeInteger, pgtypeUUID,
// enumBinding, etc.) and converts their GoTypeBinding output into the
// backend-agnostic ir.TypeBinding.
type GoTypeRegistry struct{}

// Bind maps a PGFamily + BindOpts to an ir.TypeBinding by delegating to
// the existing Go-specific binding functions and converting the result.
func (r *GoTypeRegistry) Bind(family typemap.PGFamily, opts ir.BindOpts) ir.TypeBinding {
	gb := bindGoType(family, opts)
	return toIRBinding(family, opts, gb)
}
