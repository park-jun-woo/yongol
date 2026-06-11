//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what postOpStatusResp — 지정 2xx 상태 응답을 가진 테스트용 POST PathItem 생성 (respProps nil이면 무본문 204류)

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// postOpStatusResp creates a PathItem with a POST operation whose response is
// declared under the given status code. When respProps is non-nil the response
// carries a JSON object schema with those top-level properties; when nil the
// response has no content (e.g. 204). Used to pin BUG-128 / Phase039: a 201
// body is consumable, a 204/no-body op is not.
func postOpStatusResp(opID string, status int, respProps map[string]*openapi3.SchemaRef) *openapi3.PathItem {
	resp := &openapi3.Response{}
	if respProps != nil {
		resp.Content = openapi3.NewContentWithJSONSchema(&openapi3.Schema{
			Type:       &openapi3.Types{"object"},
			Properties: respProps,
		})
	}
	op := &openapi3.Operation{
		OperationID: opID,
		Responses: openapi3.NewResponses(
			openapi3.WithStatus(status, &openapi3.ResponseRef{Value: resp}),
		),
	}
	return &openapi3.PathItem{Post: op}
}
