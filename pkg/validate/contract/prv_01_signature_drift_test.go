//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestPRV01SignatureDrift — preserved 파일 목록의 signature drift 오케스트레이션 검증

package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPRV01SignatureDrift(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nimport \"context\"\n\n"+
			"func (server *Server) ActivateWorkflow(ctx context.Context, request api.ActivateWorkflowRequestObject) (api.ActivateWorkflowResponseObject, error) { return nil, nil }\n")
	// Ground without the opID → removed-opID PRV-01 expected.
	diags := prv01SignatureDrift(buildFSWithOp(), []string{p})
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[PRV-01]") {
		t.Fatalf("expected one PRV-01 diag, got %+v", diags)
	}

	t.Run("empty paths → no diags", func(t *testing.T) {
		if d := prv01SignatureDrift(buildFSWithOp(), nil); len(d) != 0 {
			t.Errorf("expected no diags, got %+v", d)
		}
	})
}
