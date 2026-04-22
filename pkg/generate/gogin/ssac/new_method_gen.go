//ff:func feature=gen-gogin type=util control=sequence
//ff:what newMethodGen — methodGen 생성 + OpenAPI 메타데이터 주입

package ssac

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// newMethodGen extracts all needed info from OpenAPI for one operation.
// wrapCalls toggles Phase009 @call span wrapping — callers pass the
// resolved manifest.backend.observability.tracing.wrap_calls flag so the
// generator can emit otel.Tracer().Start wrappers only when explicitly
// opted-in.
func newMethodGen(doc *openapi3.T, sf ssacparser.ServiceFunc, modulePath string, useTx bool, projectFuncs, builtinFuncs []funcspec.FuncSpec, wrapCalls bool) *methodGen {
	g := &methodGen{
		FuncName:     sf.Name,
		FileName:     sf.FileName,
		ModulePath:   modulePath,
		PathParams:   make(map[string]bool),
		QueryParams:  make(map[string]queryParam),
		BodyFormats:  make(map[string]string),
		RespFields:   make(map[string]responseField),
		UseTx:        useTx,
		FirstErr:     !useTx, // tx가 있으면 이미 err 선언됨
		ProjectFuncs: projectFuncs,
		BuiltinFuncs: builtinFuncs,
		WrapCalls:    wrapCalls,
	}
	if doc != nil {
		g.extractFromOpenAPI(doc, sf.Name)
	}
	return g
}
