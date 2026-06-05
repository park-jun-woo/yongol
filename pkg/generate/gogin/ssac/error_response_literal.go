//ff:func feature=gen-gogin type=util control=selection
//ff:what errorResponseLiteral — 에러 응답 래퍼 리터럴 생성 (alias/embedded 두 형태 일원화)

package ssac

import "fmt"

// errorResponseLiteral builds the Go composite literal for an error response
// wrapper, choosing the shape from the classified ResponseShapes map (BUG-106 /
// Phase012):
//
//	alias:    api.<Op><Status>JSONResponse{Error: <msg>, Code: <code>}
//	embedded: api.<Op><Status>JSONResponse{<Embedded>: api.<Embedded>{Error: <msg>, Code: <code>}}
//
// When the wrapper type is absent from the map (or classified as alias) the
// alias form is emitted, preserving the pre-Phase012 behaviour.
func (g *methodGen) errorResponseLiteral(status int, msg, code string) string {
	typeName := fmt.Sprintf("%s%dJSONResponse", g.FuncName, status)
	if shape, ok := g.ResponseShapes[typeName]; ok && shape.Kind == shapeEmbedded {
		return fmt.Sprintf("api.%s{%s: api.%s{Error: %q, Code: %q}}",
			typeName, shape.EmbeddedType, shape.EmbeddedType, msg, code)
	}
	return fmt.Sprintf("api.%s{Error: %q, Code: %q}", typeName, msg, code)
}
