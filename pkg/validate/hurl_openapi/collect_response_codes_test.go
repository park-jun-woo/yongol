//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what collectResponseCodes — operation 응답 코드 키 집합 생성 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectResponseCodes(t *testing.T) {
	cases := []struct {
		name      string
		op        *openapi3.Operation
		wantCodes map[string]bool
	}{
		{
			name:      "nil_op",
			op:        nil,
			wantCodes: map[string]bool{},
		},
		{
			name:      "nil_responses",
			op:        &openapi3.Operation{},
			wantCodes: map[string]bool{},
		},
		{
			name: "single_response_code",
			op: func() *openapi3.Operation {
				op := &openapi3.Operation{Responses: &openapi3.Responses{}}
				op.Responses.Set("200", &openapi3.ResponseRef{Value: openapi3.NewResponse()})
				return op
			}(),
			wantCodes: map[string]bool{"200": true},
		},
		{
			name: "multiple_response_codes",
			op: func() *openapi3.Operation {
				op := &openapi3.Operation{Responses: &openapi3.Responses{}}
				op.Responses.Set("200", &openapi3.ResponseRef{Value: openapi3.NewResponse()})
				op.Responses.Set("404", &openapi3.ResponseRef{Value: openapi3.NewResponse()})
				op.Responses.Set("500", &openapi3.ResponseRef{Value: openapi3.NewResponse()})
				return op
			}(),
			wantCodes: map[string]bool{"200": true, "404": true, "500": true},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runBoolMapCase(t, collectResponseCodes(c.op), c.wantCodes)
		})
	}
}
