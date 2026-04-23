//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.extractRespFields — 200 응답 스키마의 필드를 RespFields에 수집

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// extractRespFields populates RespFields from the operation's success
// response schema, marking fields listed in `required` so buildResponse
// can skip pointer wrapping (oapi-codegen emits those as non-pointer
// Go fields). The status is picked from g.SuccessStatus, which was
// derived from the HTTP method + declared 2xx set (BUG-004), so POST
// handlers read the 201 schema rather than the non-existent 200.
func (g *methodGen) extractRespFields(op *openapi3.Operation) {
	resp := op.Responses.Status(g.SuccessStatus)
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
