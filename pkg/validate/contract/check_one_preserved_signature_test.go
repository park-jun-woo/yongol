//ff:func feature=validate-contract type=test control=sequence
//ff:what TestCheckOnePreservedSignature — preserved 파일 1건 signature drift 단일 Diagnostic 검증

package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckOnePreservedSignature(t *testing.T) {
	dir := t.TempDir()

	t.Run("unextractable signature → nil", func(t *testing.T) {
		// File with no top-level FuncDecl: ExtractSignature yields an
		// empty name, so checkOnePreservedSignature short-circuits to nil
		// rather than emitting a removed-opID diagnostic.
		p := filepath.Join(dir, "helpers.go")
		writePreserved(t, p, "package service\n\nvar X = 1\n")
		if d := checkOnePreservedSignature(buildFSWithOp().Ground(), p); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("removed opID → PRV-01", func(t *testing.T) {
		p := filepath.Join(dir, "activate_workflow.go")
		writePreserved(t, p,
			"package service\n\nimport \"context\"\n\n"+
				"func (server *Server) ActivateWorkflow(ctx context.Context, request api.ActivateWorkflowRequestObject) (api.ActivateWorkflowResponseObject, error) { return nil, nil }\n")
		diags := checkOnePreservedSignature(buildFSWithOp().Ground(), p)
		if len(diags) != 1 || !strings.Contains(diags[0].Message, "[PRV-01]") {
			t.Fatalf("expected PRV-01 diag, got %+v", diags)
		}
	})

	t.Run("matching signature → nil", func(t *testing.T) {
		p := filepath.Join(dir, "activate_workflow2.go")
		writePreserved(t, p,
			"package service\n\nimport \"context\"\n\n"+
				"func (server *Server) ActivateWorkflow2(ctx context.Context, request api.ActivateWorkflow2RequestObject) (api.ActivateWorkflow2ResponseObject, error) { return nil, nil }\n")
		if d := checkOnePreservedSignature(buildFSWithOp("ActivateWorkflow2").Ground(), p); d != nil {
			t.Errorf("matching signature should be clean, got %+v", d)
		}
	})
}
