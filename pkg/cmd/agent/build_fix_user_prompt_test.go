//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestBuildFixUserPrompt — 레이어별 cross-SSOT 컨텍스트/공통 섹션 구성 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildFixUserPrompt(t *testing.T) {
	specsDir := t.TempDir()
	ff := &features.FeaturesFile{
		Features: []features.Feature{
			{Op: "CreateWorkflow", Path: "/v1/workflows", Desc: "create a workflow", Table: "workflows"},
			// op matches the SSaC filename base so the lookup hit branch is taken.
			{Op: "create_workflow", Path: "/v1/workflows", Desc: "create", Table: "workflows"},
		},
		Tables: map[string]features.TableDef{
			"workflows": {States: []string{"draft", "active"}},
		},
	}
	diags := []diagnostic.Diagnostic{{Message: "X-1: bad", Advice: "do better"}}

	// Every layer must always emit the common "Current file" + diag + fix sections.
	assertCommon := func(t *testing.T, got, filename string) {
		t.Helper()
		if !strings.Contains(got, "Current file (") {
			t.Errorf("missing Current file section, got:\n%s", got)
		}
		if !strings.Contains(got, "X-1: bad") {
			t.Errorf("missing diagnostics, got:\n%s", got)
		}
		if !strings.Contains(got, "Fix the file.") {
			t.Errorf("missing fix instruction, got:\n%s", got)
		}
	}

	cases := []struct {
		name     string
		filename string
		l        layer
	}{
		{"DDL", "db/workflows.sql", layerDDL},
		{"SQLcQuery", "db/queries/workflows.sql", layerSQLcQuery},
		{"SSaC", "service/workflow/create_workflow.ssac", layerSSaC},
		{"StateDiagram", "states/workflows.mmd", layerStateDiagram},
		{"OpenAPIDefault", "api/openapi.yaml", layerOpenAPI},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := buildFixUserPrompt(specsDir, ff, tc.filename, "FILE_CONTENT", diags, tc.l)
			assertCommon(t, got, tc.filename)
			if !strings.Contains(got, "FILE_CONTENT") {
				t.Errorf("expected file content embedded, got:\n%s", got)
			}
		})
	}
}
