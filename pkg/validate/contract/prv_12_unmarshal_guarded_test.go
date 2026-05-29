//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV12UnmarshalGuarded — err 체크 있는 Unmarshal 은 ERROR 없음

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV12UnmarshalGuarded(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nimport \"encoding/json\"\n\n"+
			"func h(body []byte, req any) error {\n"+
			"\tif err := json.Unmarshal(body, req); err != nil { return err }\n"+
			"\treturn nil\n}\n")
	diags := prv12PreservedUnmarshalErr([]string{p})
	if len(diags) != 0 {
		t.Fatalf("guarded Unmarshal should be safe, got %+v", diags)
	}
}
