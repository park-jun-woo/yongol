//ff:func feature=rule type=test control=iteration dimension=1
//ff:what registerParamGoType test — 파라미터 schema → OpenAPI.paramType.<op>.<name> 등록 (맥락·array·skip)

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// strParam / arrayParam / refParam build parameters whose Schema mirrors the
// shapes registerParamGoType must resolve in the CtxParam context.
func strParam(name, format string) *openapi3.Parameter {
	return &openapi3.Parameter{Name: name, In: "query", Schema: strSchema(format)}
}

// TestRegisterParamGoType verifies the Go type registered under
// "OpenAPI.paramType.<opID>.<name>" — the key whose sole type-value consumer is
// xfs_73. It pins (a) the CtxParam context (date-time param → string, NOT the
// response-body time.Time), (b) the array/number blind-spot closure (array and
// number params now register as []T / float32 instead of being skipped), and
// (c) the skip rule (a schema resolving to "" registers nothing, leaving the
// key absent — exercised by the nil-schema case).
func TestRegisterParamGoType(t *testing.T) {
	const op = "Op"
	const key = "OpenAPI.paramType.Op."

	tests := []struct {
		name      string
		param     *openapi3.Parameter
		wantType  string // expected registered value; "" means key must be ABSENT
		wantFound bool
	}{
		// scalar string formats — uuid/email context-independent
		{"uuid param", strParam("u", "uuid"), "openapi_types.UUID", true},
		{"email param", strParam("e", "email"), "openapi_types.Email", true},
		{"plain string param", strParam("s", ""), "string", true},
		// CONTEXT divergence: a date-time *parameter* is plain string,
		// NOT time.Time — this kills a CtxParam→CtxResponseBody mutation.
		{"date-time param stays string", strParam("dt", "date-time"), "string", true},
		// formatless integer param is int (context-independent, matches response)
		{"plain integer param int", &openapi3.Parameter{
			Name: "n", In: "query", Schema: intSchema("")}, "int", true},
		// array blind-spot closure: array-uuid param registers as []T,
		// formerly skipped entirely.
		{"array-uuid param []T", &openapi3.Parameter{
			Name: "ids", In: "query", Schema: arraySchema(strSchema("uuid"))},
			"[]openapi_types.UUID", true},
		// $ref param registers under the referenced type name
		{"ref param", &openapi3.Parameter{
			Name: "wf", In: "query", Schema: refSchema("Workflow")}, "Workflow", true},
		// SKIP rule: a nil schema resolves to "" → nothing registered.
		{"nil schema not registered", &openapi3.Parameter{Name: "x", In: "query"}, "", false},
		// number param registers as float32 (formerly skipped — array/number blind
		// spot closure; oapi-codegen renders a formatless/float number param as float32).
		{"number param registers float32", &openapi3.Parameter{
			Name: "y", In: "query", Schema: numSchema("float")}, "float32", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGround()
			registerParamGoType(g, op, tt.param)
			got, found := g.Types[key+tt.param.Name]
			if found != tt.wantFound {
				t.Fatalf("%s: key present=%v, want %v (value=%q)", tt.name, found, tt.wantFound, got)
			}
			if found && got != tt.wantType {
				t.Errorf("%s: registered %q, want %q", tt.name, got, tt.wantType)
			}
		})
	}
}
