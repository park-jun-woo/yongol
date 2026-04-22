//ff:func feature=validate-contract type=test control=sequence
//ff:what TestPRV01RemovedOpID — operationId 가 SSOT 에서 사라지면 PRV-01 ERROR 발행

package contract

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestPRV01RemovedOpID(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nimport \"context\"\n\n"+
			"func (server *Server) ActivateWorkflow(ctx context.Context, request api.ActivateWorkflowRequestObject) (api.ActivateWorkflowResponseObject, error) { return nil, nil }\n")
	fs := buildFSWithOp()
	diags := prv01SignatureDrift(fs, []string{p})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d", len(diags))
	}
	if diags[0].Level != diagnostic.LevelError {
		t.Errorf("expected ERROR level")
	}
	if !strings.Contains(diags[0].Message, "[PRV-01]") {
		t.Errorf("expected [PRV-01] prefix in message: %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "ActivateWorkflow") {
		t.Errorf("expected opID in message: %q", diags[0].Message)
	}
}
