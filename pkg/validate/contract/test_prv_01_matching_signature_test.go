//ff:func feature=validate-contract type=test control=sequence
//ff:what TestPRV01MatchingSignatureNoDiag — signature 일치 시 PRV-01 ERROR 없음 검증

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV01MatchingSignatureNoDiag(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nimport \"context\"\n\n"+
			"func (server *Server) ActivateWorkflow(ctx context.Context, request api.ActivateWorkflowRequestObject) (api.ActivateWorkflowResponseObject, error) { return nil, nil }\n")
	fs := buildFSWithOp("ActivateWorkflow")
	diags := prv01SignatureDrift(fs, []string{p})
	if len(diags) != 0 {
		t.Errorf("expected no diagnostics, got %+v", diags)
	}
}
