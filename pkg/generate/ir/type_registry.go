//ff:type feature=gen-ir type=model
//ff:what TypeRegistry — PGFamily + BindOpts → TypeBinding 변환 인터페이스 (백엔드별 구현)

package ir

import "github.com/park-jun-woo/yongol/pkg/generate/typemap"

// TypeRegistry converts a PGFamily + BindOpts pair into a backend-specific
// TypeBinding. Each backend target (Go+Gin, NestJS, FastAPI) provides its
// own implementation.
type TypeRegistry interface {
	// Bind produces the TypeBinding for the given family and options.
	Bind(family typemap.PGFamily, opts BindOpts) TypeBinding
}
