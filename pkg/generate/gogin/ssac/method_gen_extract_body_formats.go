//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.extractBodyFormats — request body 스키마에서 format별 wrapper cast 정보 수집

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// extractBodyFormats populates BodyFormats from the request body schema so
// mapValue can emit an explicit primitive cast (e.g. string(request.Body.Email))
// when the OpenAPI format produces an oapi-codegen wrapper type.
func (g *methodGen) extractBodyFormats(op *openapi3.Operation) {
	if op.RequestBody == nil || op.RequestBody.Value == nil {
		return
	}
	mt := op.RequestBody.Value.Content.Get("application/json")
	if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
		return
	}
	for propName, propRef := range mt.Schema.Value.Properties {
		if propRef == nil || propRef.Value == nil {
			continue
		}
		if propRef.Value.Format != "" {
			g.BodyFormats[propName] = propRef.Value.Format
		}
	}
}
