//ff:type feature=gen-nestjs type=adapter
//ff:what NestJSTypeRegistry — ir.TypeRegistry 의 NestJS 구현 (PGFamily + BindOpts → TypeScript/Prisma TypeBinding)

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// NestJSTypeRegistry implements ir.TypeRegistry for the NestJS (TypeScript
// + Prisma) backend. Each PGFamily maps to a Prisma schema type (DBType),
// a TypeScript API type (APIType), and the conversion expressions between
// DB and API layers.
type NestJSTypeRegistry struct{}

// Bind maps a PGFamily + BindOpts to an ir.TypeBinding containing
// Prisma schema type, TypeScript API type, and TS↔DB conversion
// expressions.
func (r *NestJSTypeRegistry) Bind(family typemap.PGFamily, opts ir.BindOpts) ir.TypeBinding {
	return bindNestJS(family, opts)
}
