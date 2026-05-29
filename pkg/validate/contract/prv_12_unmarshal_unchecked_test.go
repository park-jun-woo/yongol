//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV12UnmarshalUnchecked — Unmarshal 에러 미체크 ERROR 발생

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV12UnmarshalUnchecked(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nimport \"encoding/json\"\n\n"+
			"func h(body []byte, req any) {\n"+
			"\terr := json.Unmarshal(body, req)\n"+
			"\t_ = err\n}\n")
	diags := prv12PreservedUnmarshalErr([]string{p})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
}
