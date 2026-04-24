//ff:func feature=validate type=test-helper control=selection topic=states
//ff:what buildXsm27Fixture — TestXsm27 용 최소 Fullstack 조립

package ssac_statemachine

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildXsm27Fixture constructs a minimal Fullstack covering one endpoint
// with a single SSaC function and (optionally) a Workflow state diagram +
// matching DDL DEFAULT so XSM-27's stateful-resource gate fires.
func buildXsm27Fixture(tc xsm27Case) *yongol.Fullstack {
	paths := openapi3.NewPaths()
	item := &openapi3.PathItem{}
	op := &openapi3.Operation{OperationID: tc.opID, Responses: openapi3.NewResponses()}
	switch tc.method {
	case "GET":
		item.Get = op
	case "POST":
		item.Post = op
	case "PUT":
		item.Put = op
	case "DELETE":
		item.Delete = op
	}
	paths.Set(tc.path, item)

	var diagrams []*statemachine.StateDiagram
	g := &rule.Ground{
		Lookup: map[string]rule.StringSet{},
		Types:  map[string]string{},
		Pairs:  map[string]rule.StringSet{},
		Config: map[string]bool{},
		Vars:   rule.StringSet{},
		Flags:  rule.StringSet{},
	}
	if !tc.nonStatefulRes {
		if tc.withDiagram {
			diagrams = append(diagrams, &statemachine.StateDiagram{
				ID:           "workflow",
				Symbol:       "Workflow",
				InitialState: "draft",
				States:       []string{"draft", "active"},
			})
		}
		if tc.withDefault {
			g.Types["DDL.default.value.workflows.status"] = "draft"
		}
	}
	fs := &yongol.Fullstack{
		OpenAPIDoc: &openapi3.T{
			OpenAPI: "3.0.0",
			Info:    &openapi3.Info{Title: "t", Version: "1"},
			Paths:   paths,
		},
		StateDiagrams: diagrams,
		ServiceFuncs: []ssac.ServiceFunc{{
			Name:         tc.opID,
			FileName:     "service/workflow/" + strings.ToLower(tc.opID) + ".ssac",
			Line:         5,
			Sequences:    tc.sequences,
			StateNeutral: tc.stateNeutral,
		}},
	}
	fs.SetGround(g)
	return fs
}
