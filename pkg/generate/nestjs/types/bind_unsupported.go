//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindUnsupported — FamilyUnsupported → Supported=false 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindUnsupported returns an ir.TypeBinding with Supported=false for
// unrecognised PG families. validate D-11 rejects these before code
// generation runs.
func bindUnsupported(family typemap.PGFamily) ir.TypeBinding {
	return ir.TypeBinding{
		Family:    family,
		Supported: false,
	}
}
