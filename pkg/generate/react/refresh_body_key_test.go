//ff:func feature=gen-react type=test control=sequence
//ff:what refreshBodyKey — requestBody JSON 에 refresh_field 프로퍼티 존재 시 키 반환, 없으면 빈 문자열

package react

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestRefreshBodyKey(t *testing.T) {
	// nil op -> ""
	if got := refreshBodyKey(nil, "refresh_token"); got != "" {
		t.Errorf("nil op = %q, want empty", got)
	}

	// op without requestBody -> "" (cookie-carried refresh)
	if got := refreshBodyKey(&openapi3.Operation{}, "refresh_token"); got != "" {
		t.Errorf("no body = %q, want empty", got)
	}

	// requestBody declares the refresh field -> field name returned
	withBody := buildTokenOp("Refresh", []string{"access_token"}, []string{"refresh_token"})
	if got := refreshBodyKey(withBody, "refresh_token"); got != "refresh_token" {
		t.Errorf("matched body = %q, want refresh_token", got)
	}

	// requestBody present but lacks the refresh field -> ""
	otherBody := buildTokenOp("Refresh", []string{"access_token"}, []string{"email"})
	if got := refreshBodyKey(otherBody, "refresh_token"); got != "" {
		t.Errorf("unmatched body = %q, want empty", got)
	}

	// non-JSON requestBody content -> "" (skipped)
	formBody := &openapi3.Operation{RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
		Content: openapi3.Content{"application/x-www-form-urlencoded": &openapi3.MediaType{
			Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
				Properties: openapi3.Schemas{"refresh_token": &openapi3.SchemaRef{Value: &openapi3.Schema{}}},
			}},
		}},
	}}}
	if got := refreshBodyKey(formBody, "refresh_token"); got != "" {
		t.Errorf("non-JSON body = %q, want empty", got)
	}

	// JSON media type but schema-less -> "" (no panic)
	schemaless := &openapi3.Operation{RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
		Content: openapi3.Content{"application/json": &openapi3.MediaType{}},
	}}}
	if got := refreshBodyKey(schemaless, "refresh_token"); got != "" {
		t.Errorf("schemaless body = %q, want empty", got)
	}
}
