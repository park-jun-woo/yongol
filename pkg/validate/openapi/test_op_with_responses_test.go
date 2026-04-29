//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=response-body-required
//ff:what 테스트 헬퍼 — operationId + status→ResponseRef 매핑으로 Operation 빌드

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// opWithResponses builds an Operation with the given operationId and
// status→ResponseRef mapping. Used by O-5 unit tests.
func opWithResponses(opID string, statusToResp map[string]*openapi3.ResponseRef) *openapi3.Operation {
	resps := openapi3.NewResponses()
	for status, ref := range statusToResp {
		resps.Set(status, ref)
	}
	return &openapi3.Operation{
		OperationID: opID,
		Responses:   resps,
	}
}
