//ff:func feature=rule type=test control=iteration dimension=2
//ff:what resolveOAPIGoType test — type×format×shape×context 매트릭스 (oapi-codegen ground truth 대조)

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// intSchema / numSchema / boolSchema / objSchema build leaf schema refs for the
// matrix. strSchema / arraySchema are defined in
// register_openapi_response_props_test.go (same package).
func intSchema(format string) *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type: &openapi3.Types{"integer"}, Format: format}}
}

func numSchema(format string) *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type: &openapi3.Types{"number"}, Format: format}}
}

func boolSchema() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"boolean"}}}
}

func objSchema() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"object"}}}
}

func refSchema(name string) *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Ref: "#/components/schemas/" + name}
}

func TestResolveOAPIGoType(t *testing.T) {
	tests := []struct {
		name string
		ref  *openapi3.SchemaRef
		ctx  OAPIContext
		want string
	}{
		// --- scalar string × format × context ---
		{"resp string plain", strSchema(""), CtxResponseBody, "string"},
		{"resp string uuid", strSchema("uuid"), CtxResponseBody, "openapi_types.UUID"},
		{"resp string email", strSchema("email"), CtxResponseBody, "openapi_types.Email"},
		{"resp string date-time", strSchema("date-time"), CtxResponseBody, "time.Time"},
		{"param string plain", strSchema(""), CtxParam, "string"},
		{"param string uuid", strSchema("uuid"), CtxParam, "openapi_types.UUID"},
		{"param string email", strSchema("email"), CtxParam, "openapi_types.Email"},
		// context divergence: param date-time → string (NOT time.Time)
		{"param string date-time", strSchema("date-time"), CtxParam, "string"},

		// --- integer × format × context ---
		{"resp int plain", intSchema(""), CtxResponseBody, "int"},
		{"resp int int64", intSchema("int64"), CtxResponseBody, "int64"},
		{"resp int int32", intSchema("int32"), CtxResponseBody, "int"},
		{"param int plain", intSchema(""), CtxParam, "int32"},
		{"param int int64", intSchema("int64"), CtxParam, "int64"},
		{"param int int32", intSchema("int32"), CtxParam, "int32"},

		// --- number × format × context ---
		{"resp number plain", numSchema(""), CtxResponseBody, "float64"},
		{"resp number float", numSchema("float"), CtxResponseBody, "float32"},
		{"param number float", numSchema("float"), CtxParam, ""},

		// --- boolean / object ---
		{"resp boolean", boolSchema(), CtxResponseBody, "bool"},
		{"param boolean", boolSchema(), CtxParam, "bool"},
		{"resp object", objSchema(), CtxResponseBody, "object"},
		{"param object", objSchema(), CtxParam, ""},

		// --- arrays (recursion, both contexts) ---
		{"resp []uuid", arraySchema(strSchema("uuid")), CtxResponseBody, "[]openapi_types.UUID"},
		{"resp []date-time", arraySchema(strSchema("date-time")), CtxResponseBody, "[]time.Time"},
		{"resp []string", arraySchema(strSchema("")), CtxResponseBody, "[]string"},
		{"resp []int64", arraySchema(intSchema("int64")), CtxResponseBody, "[]int64"},
		// array-uuid PARAMETER — the formerly-skipped blind spot, now []openapi_types.UUID
		{"param []uuid", arraySchema(strSchema("uuid")), CtxParam, "[]openapi_types.UUID"},
		{"param []string", arraySchema(strSchema("")), CtxParam, "[]string"},
		// param date-time array stays []string (context honoured through recursion)
		{"param []date-time", arraySchema(strSchema("date-time")), CtxParam, "[]string"},

		// --- nested arrays (recursion, no per-shape branch) ---
		{"resp [][]uuid", arraySchema(arraySchema(strSchema("uuid"))), CtxResponseBody, "[][]openapi_types.UUID"},
		{"param [][]uuid", arraySchema(arraySchema(strSchema("uuid"))), CtxParam, "[][]openapi_types.UUID"},

		// --- $ref → type name (both contexts) ---
		{"resp ref", refSchema("Workflow"), CtxResponseBody, "Workflow"},
		{"param ref", refSchema("Workflow"), CtxParam, "Workflow"},

		// --- edge cases ---
		{"nil ref", nil, CtxResponseBody, ""},
		{"untyped value", &openapi3.SchemaRef{Value: &openapi3.Schema{}}, CtxResponseBody, ""},
		{"array nil items", &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"array"}}}, CtxResponseBody, ""},
		// unsupported JSON type (e.g. null) falls through every case → ""
		{"unsupported type", &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"null"}}}, CtxResponseBody, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveOAPIGoType(tt.ref, tt.ctx)
			if got != tt.want {
				t.Errorf("resolveOAPIGoType(%s) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
