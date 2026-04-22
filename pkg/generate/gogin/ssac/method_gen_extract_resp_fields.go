//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.extractRespFields — 200 응답 스키마의 필드를 RespFields에 수집

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// extractRespFields populates RespFields from the 200 response schema,
// marking fields listed in `required` so buildResponse can skip pointer
// wrapping (oapi-codegen emits those as non-pointer Go fields).
func (g *methodGen) extractRespFields(op *openapi3.Operation) {
	resp := op.Responses.Status(200)
	if resp == nil || resp.Value == nil || resp.Value.Content == nil {
		return
	}
	mt := resp.Value.Content.Get("application/json")
	if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
		return
	}
	required := make(map[string]bool, len(mt.Schema.Value.Required))
	for _, name := range mt.Schema.Value.Required {
		required[name] = true
	}
	for propName, propRef := range mt.Schema.Value.Properties {
		g.RespFields[propName] = buildResponseField(propName, propRef, required[propName])
	}
}
