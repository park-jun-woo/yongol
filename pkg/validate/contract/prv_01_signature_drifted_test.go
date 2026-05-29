//ff:func feature=validate-contract type=test control=sequence
//ff:what TestPRV01SignatureDriftParamType — signature drift(파라미터 축소) 감지

package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPRV01SignatureDriftParamType(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nimport \"context\"\n\n"+
			"func (server *Server) ActivateWorkflow(ctx context.Context) error { return nil }\n")
	fs := buildFSWithOp("ActivateWorkflow")
	diags := prv01SignatureDrift(fs, []string{p})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "signature drifted") {
		t.Errorf("expected drift message, got %q", diags[0].Message)
	}
}
