//ff:func feature=rule type=test control=iteration dimension=1
//ff:what registerParamGoType test — 파라미터 schema → OpenAPI.paramType.<op>.<name> 등록 (맥락·array·skip)
package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestRegisterParamGoType(t *testing.T) {
	const op = "Op"

	tests := []struct {
		name      string
		param     *openapi3.Parameter
		wantType  string // expected registered value; "" means key must be ABSENT
		wantFound bool
	}{
		{"uuid param", strParam("u", "uuid"), "openapi_types.UUID", true},
		{"email param", strParam("e", "email"), "openapi_types.Email", true},
		{"plain string param", strParam("s", ""), "string", true},
		{"date-time param stays string", strParam("dt", "date-time"), "string", true},
		{"plain integer param int", &openapi3.Parameter{
			Name: "n", In: "query", Schema: intSchema("")}, "int", true},
		{"array-uuid param []T", &openapi3.Parameter{
			Name: "ids", In: "query", Schema: arraySchema(strSchema("uuid"))},
			"[]openapi_types.UUID", true},
		{"ref param", &openapi3.Parameter{
			Name: "wf", In: "query", Schema: refSchema("Workflow")}, "Workflow", true},
		{"nil schema not registered", &openapi3.Parameter{Name: "x", In: "query"}, "", false},
		{"number param registers float32", &openapi3.Parameter{
			Name: "y", In: "query", Schema: numSchema("float")}, "float32", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertRegisterParamGoType(t, op, tt.name, tt.param, tt.wantType, tt.wantFound)
		})
	}
}
