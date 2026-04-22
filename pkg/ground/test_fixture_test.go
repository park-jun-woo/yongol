//ff:func feature=rule type=test-helper control=sequence
//ff:what 테스트용 최소 Fullstack/Ground 조립 헬퍼 — populate_* 단위 테스트 공통 fixture

package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// newGround returns a fully initialized *rule.Ground with empty maps so
// populate_* functions can write without nil-panics.
func newGround() *rule.Ground {
	return &rule.Ground{
		Lookup:  make(map[string]rule.StringSet),
		Types:   make(map[string]string),
		Pairs:   make(map[string]rule.StringSet),
		Config:  make(map[string]bool),
		Vars:    make(rule.StringSet),
		Flags:   make(rule.StringSet),
		Schemas: make(map[string][]string),
	}
}

// newMinimalFullstack returns an empty *yongol.Fullstack with nil SSOT fields.
// Opts mutate individual fields.
func newMinimalFullstack(opts ...func(*yongol.Fullstack)) *yongol.Fullstack {
	fs := &yongol.Fullstack{}
	for _, o := range opts {
		o(fs)
	}
	return fs
}

// withOpenAPIDoc attaches a prebuilt *openapi3.T.
func withOpenAPIDoc(doc *openapi3.T) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.OpenAPIDoc = doc }
}

// withDDLTables attaches DDL tables.
func withDDLTables(tables ...ddl.Table) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.DDLTables = append(fs.DDLTables, tables...) }
}

// withServiceFuncs attaches parsed SSaC service funcs.
func withServiceFuncs(funcs ...ssac.ServiceFunc) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.ServiceFuncs = append(fs.ServiceFuncs, funcs...) }
}

// withParsedPolicies attaches Rego policies.
func withParsedPolicies(pols ...rego.Policy) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.ParsedPolicies = append(fs.ParsedPolicies, pols...) }
}

// withStateDiagrams attaches Mermaid state diagrams.
func withStateDiagrams(sds ...*statemachine.StateDiagram) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.StateDiagrams = append(fs.StateDiagrams, sds...) }
}

// withRequestConstraints attaches pre-built request constraint maps.
func withRequestConstraints(m map[string]map[string]oapiparser.FieldConstraint) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.RequestConstraints = m }
}

// withResponseConstraints attaches pre-built response constraint maps.
func withResponseConstraints(m map[string]map[string]oapiparser.FieldConstraint) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.ResponseConstraints = m }
}

// intPtr returns a pointer to an int (handy for *FieldConstraint.MaxLength).
func intPtr(n int) *int { return &n }

// buildDocWithOp returns a minimal *openapi3.T containing a single path with a
// single operation, optionally with a 2xx JSON response whose schema lists the
// given top-level field names as strings.
func buildDocWithOp(path, method, opID string, respFields []string) *openapi3.T {
	op := &openapi3.Operation{OperationID: opID}
	if respFields != nil {
		props := openapi3.Schemas{}
		for _, f := range respFields {
			props[f] = &openapi3.SchemaRef{Value: &openapi3.Schema{
				Type: &openapi3.Types{"string"},
			}}
		}
		resp := openapi3.NewResponse().
			WithContent(openapi3.NewContentWithJSONSchema(&openapi3.Schema{
				Type:       &openapi3.Types{"object"},
				Properties: props,
			}))
		responses := openapi3.NewResponses()
		responses.Set("200", &openapi3.ResponseRef{Value: resp})
		op.Responses = responses
	}
	item := &openapi3.PathItem{}
	switch method {
	case "GET":
		item.Get = op
	case "POST":
		item.Post = op
	case "PUT":
		item.Put = op
	case "DELETE":
		item.Delete = op
	}
	return &openapi3.T{Paths: openapi3.NewPaths(openapi3.WithPath(path, item))}
}

// setJSONResponse overwrites the operation's response map with the given code
// and a JSON schema whose properties are the given string fields.
func setJSONResponse(op *openapi3.Operation, code string, fields []string) {
	props := openapi3.Schemas{}
	for _, f := range fields {
		props[f] = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	}
	resp := openapi3.NewResponse().WithContent(openapi3.NewContentWithJSONSchema(&openapi3.Schema{
		Type:       &openapi3.Types{"object"},
		Properties: props,
	}))
	if op.Responses == nil {
		op.Responses = openapi3.NewResponses()
	}
	op.Responses.Set(code, &openapi3.ResponseRef{Value: resp})
}
