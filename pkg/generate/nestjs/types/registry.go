//ff:type feature=gen-nestjs type=adapter
//ff:what NestJSTypeRegistry — ir.TypeRegistry 의 NestJS 스켈레톤 구현 (미구현: Supported=false)

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// NestJSTypeRegistry implements ir.TypeRegistry for the NestJS (TypeScript
// + Prisma) backend. This is a skeleton — all families return
// Supported=false until the NestJS backend is fully implemented.
type NestJSTypeRegistry struct{}

// Bind always returns an unsupported TypeBinding. The NestJS backend type
// mappings (TypeScript type names, Prisma schema types, TS↔DB conversion
// expressions) will be filled in when the NestJS code generator is built.
func (r *NestJSTypeRegistry) Bind(family typemap.PGFamily, opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:    family,
		NotNull:   opts.NotNull,
		DBType:    "/* nestjs: not yet implemented */",
		APIType:   "/* nestjs: not yet implemented */",
		Supported: false,
	}
}
