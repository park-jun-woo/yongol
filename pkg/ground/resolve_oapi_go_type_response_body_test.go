//ff:func feature=rule type=test control=iteration dimension=1
//ff:what resolveOAPIGoType(CtxResponseBody) test — 응답 본문 필드 type+format → Go 타입 (array items format-aware, BUG-102)
package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestResolveOAPIGoTypeResponseBody(t *testing.T) {
	tests := []struct {
		name string
		ref  *openapi3.SchemaRef
		want string
	}{
		// top-level string (regression — must stay format-aware)
		{"string uuid", strSchema("uuid"), "openapi_types.UUID"},
		{"string date-time", strSchema("date-time"), "time.Time"},
		{"string email", strSchema("email"), "openapi_types.Email"},
		{"string plain", strSchema(""), "string"},

		// array items — the BUG-102 fix: items format must be honoured
		{"array of uuid", arraySchema(strSchema("uuid")), "[]openapi_types.UUID"},
		{"array of date-time", arraySchema(strSchema("date-time")), "[]time.Time"},
		{"array of email", arraySchema(strSchema("email")), "[]openapi_types.Email"},
		{"array of plain string", arraySchema(strSchema("")), "[]string"},

		// non-string array items keep the resolveSchemaType path
		{"array of int64", arraySchema(&openapi3.SchemaRef{Value: &openapi3.Schema{
			Type: &openapi3.Types{"integer"}, Format: "int64"}}), "[]int64"},

		// $ref / nil fall through to resolveSchemaType
		{"ref", &openapi3.SchemaRef{Ref: "#/components/schemas/Workflow"}, "Workflow"},
		{"nil", nil, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveOAPIGoType(tt.ref, CtxResponseBody)
			if got != tt.want {
				t.Errorf("resolveOAPIGoType(%s, CtxResponseBody) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
