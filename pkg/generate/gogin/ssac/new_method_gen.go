//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what newMethodGen — methodGen 생성 + OpenAPI 메타데이터 주입

package ssac

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// newMethodGen extracts all needed info from OpenAPI for one operation.
// wrapCalls toggles Phase009 @call span wrapping — callers pass the
// resolved manifest.backend.observability.tracing.wrap_calls flag so the
// generator can emit otel.Tracer().Start wrappers only when explicitly
// opted-in.
//
// diagrams supplies the parsed Mermaid stateDiagrams; newMethodGen
// flattens them into a `DiagramID → Symbol` lookup so build_state can
// emit `statemachine.<Symbol>CanTransition(...)` even when the source
// .md file was authored in lowercase (BUG-002).
func newMethodGen(doc *openapi3.T, sf ssacparser.ServiceFunc, modulePath string, useTx bool, projectFuncs, builtinFuncs []funcspec.FuncSpec, wrapCalls bool, diagrams []*statemachine.StateDiagram, ownerships []rego.OwnershipMapping, ddlTables []ddl.Table, sqlcQueries []sqlcparser.QuerySpec) *methodGen {
	symbols := make(map[string]string, len(diagrams))
	for _, d := range diagrams {
		if d == nil {
			continue
		}
		symbols[d.ID] = d.Symbol
	}
	// Pre-compute the var→type map from all Result bindings in the
	// service func's sequences. Populated before any buildX call so a
	// later @response can look up the model even when the assignment
	// appeared earlier in the sequence list.
	varTypes := make(map[string]string)
	for _, seq := range sf.Sequences {
		if seq.Result != nil && seq.Result.Var != "" && seq.Result.Type != "" {
			varTypes[seq.Result.Var] = seq.Result.Type
		}
	}
	g := &methodGen{
		SuccessStatus: 200, // overwritten by extractFromOpenAPI for real ops
		VarTypes:        varTypes,
		BodyJSONBFields: make(map[string]bool),
		FuncName:        sf.Name,
		FileName:      sf.FileName,
		ModulePath:    modulePath,
		PathParams:    make(map[string]bool),
		QueryParams:   make(map[string]queryParam),
		BodyFormats:   make(map[string]string),
		RespFields:    make(map[string]responseField),
		UseTx:         useTx,
		FirstErr:      !useTx, // tx가 있으면 이미 err 선언됨
		ProjectFuncs:  projectFuncs,
		BuiltinFuncs:  builtinFuncs,
		WrapCalls:     wrapCalls,
		DiagramSymbol: symbols,
		Ownerships:    ownerships,
		DDLTables:     ddlTables,
		SQLcQueries:   sqlcQueries,
	}
	if doc != nil {
		g.extractFromOpenAPI(doc, sf.Name)
	}
	return g
}
