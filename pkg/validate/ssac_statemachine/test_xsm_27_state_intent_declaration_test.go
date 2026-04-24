//ff:func feature=validate type=test control=iteration dimension=1 topic=states
//ff:what XSM-27 테이블 기반 테스트 — state intent declaration (@state / @state-neutral) 강제

package ssac_statemachine

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildXsm27Fixture constructs a minimal Fullstack with one OpenAPI
// operation, one SSaC function, and (optionally) a workflow stateDiagram +
// DDL DEFAULT. A caller-visible builder keeps each test case compact.
type xsm27Case struct {
	name           string
	method         string                 // "GET" / "POST" / "PUT" / "DELETE"
	path           string                 // e.g. "/workflows/{id}/execute"
	opID           string
	sequences      []ssac.Sequence
	stateNeutral   bool
	withDiagram    bool
	withDefault    bool
	nonStatefulRes bool                   // when true skip diagram/default → path resource not stateful
	wantFire       bool
}

func TestXsm27StateIntentDeclaration(t *testing.T) {
	findByID := ssac.Sequence{
		Type:  "get",
		Model: "Workflow.FindByID",
		Args: []ssac.Arg{
			{Source: "request", Field: "id"},
		},
		Result: &ssac.Result{Type: "Workflow", Var: "wf"},
	}
	stateSeq := ssac.Sequence{
		Type:       "state",
		DiagramID:  "workflow",
		Inputs:     map[string]string{"Status": "wf.Status"},
		Transition: "ActivateWorkflow",
		Message:    "Cannot activate",
	}
	nonStatefulGet := ssac.Sequence{
		Type:  "get",
		Model: "Comment.FindByID",
		Args: []ssac.Arg{
			{Source: "request", Field: "id"},
		},
		Result: &ssac.Result{Type: "Comment", Var: "c"},
	}

	cases := []xsm27Case{
		{
			name:        "ExecuteWorkflow (POST, stateful, no state, no neutral) fires",
			method:      "POST",
			path:        "/workflows/{id}/execute",
			opID:        "ExecuteWorkflow",
			sequences:   []ssac.Sequence{findByID},
			withDiagram: true,
			withDefault: true,
			wantFire:    true,
		},
		{
			name:        "ActivateWorkflow (POST, with @state) passes",
			method:      "POST",
			path:        "/workflows/{id}/activate",
			opID:        "ActivateWorkflow",
			sequences:   []ssac.Sequence{findByID, stateSeq},
			withDiagram: true,
			withDefault: true,
			wantFire:    false,
		},
		{
			name:         "LikeWorkflow (POST, with @state-neutral) passes",
			method:       "POST",
			path:         "/workflows/{id}/like",
			opID:         "LikeWorkflow",
			sequences:    []ssac.Sequence{findByID},
			stateNeutral: true,
			withDiagram:  true,
			withDefault:  true,
			wantFire:     false,
		},
		{
			name:        "ListWorkflows (GET) — method mismatch, passes",
			method:      "GET",
			path:        "/workflows/{id}",
			opID:        "GetWorkflow",
			sequences:   []ssac.Sequence{findByID},
			withDiagram: true,
			withDefault: true,
			wantFire:    false,
		},
		{
			name:        "CreateWorkflow (POST, no {id}) — passes",
			method:      "POST",
			path:        "/workflows",
			opID:        "CreateWorkflow",
			sequences:   []ssac.Sequence{},
			withDiagram: true,
			withDefault: true,
			wantFire:    false,
		},
		{
			name:           "DeleteComment (POST on non-stateful resource) — passes",
			method:         "DELETE",
			path:           "/comments/{id}",
			opID:           "DeleteComment",
			sequences:      []ssac.Sequence{nonStatefulGet},
			nonStatefulRes: true,
			wantFire:       false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			fs := buildXsm27Fixture(tc)
			diags := xsm27StateIntentDeclaration(fs)
			got := len(diags) > 0
			if got != tc.wantFire {
				t.Fatalf("want fire=%v, got %d diagnostics: %+v", tc.wantFire, len(diags), diags)
			}
			if tc.wantFire {
				d := diags[0]
				if !strings.Contains(d.Message, "[XSM-27]") {
					t.Errorf("expected [XSM-27] in message, got %q", d.Message)
				}
				if !strings.Contains(d.Advice, "Option A") || !strings.Contains(d.Advice, "Option B") {
					t.Errorf("advice must carry Option A and Option B, got %q", d.Advice)
				}
				if !strings.Contains(d.Advice, "@state-neutral") {
					t.Errorf("advice must mention @state-neutral, got %q", d.Advice)
				}
			}
		})
	}
}

// buildXsm27Fixture constructs a minimal Fullstack covering one endpoint.
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
