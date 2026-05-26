//ff:func feature=gen-nestjs type=util control=sequence
//ff:what bindBytea — FamilyBytea → Prisma Bytes / TypeScript Buffer 바인딩

package types

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
)

// bindBytea produces the TypeBinding for BYTEA columns. Prisma: Bytes,
// TypeScript: Buffer. DB insert wraps with Buffer.from(). The Buffer
// import from the Node.js "buffer" module is included in DBImports.
func bindBytea(opts ir.BindOpts) ir.TypeBinding {
	return ir.TypeBinding{
		Family:         typemap.FamilyBytea,
		NotNull:        opts.NotNull,
		DBType:         "Bytes",
		DBImports:      []string{"Buffer from buffer"},
		APIType:        nullableAPIType("Buffer", opts.NotNull),
		ToDBExpr:       "Buffer.from({var})",
		ToAPIExpr:      "{row}.{field}",
		ToResponseExpr: "{var}",
		NilCheckExpr:   nilCheckExpr(opts.NotNull),
		Supported:      true,
		DefaultLiteral: opts.DefaultLiteral,
	}
}
