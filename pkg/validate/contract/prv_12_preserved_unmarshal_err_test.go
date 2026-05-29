//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestPRV12PreservedUnmarshalErr — Unmarshal 에러 무시 오케스트레이션 검증

package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPRV12PreservedUnmarshalErr(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "h.go")
	writePreserved(t, p,
		"package service\nfunc F(b []byte) {\n"+
			"  var v T\n"+
			"  json.Unmarshal(b, &v)\n"+
			"  _ = v\n}\n")
	diags := prv12PreservedUnmarshalErr([]string{p})
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[PRV-12]") {
		t.Fatalf("expected one PRV-12 diag, got %+v", diags)
	}

	t.Run("guarded is safe", func(t *testing.T) {
		q := filepath.Join(dir, "safe.go")
		writePreserved(t, q,
			"package service\nfunc G(b []byte) error {\n"+
				"  var v T\n"+
				"  if err := json.Unmarshal(b, &v); err != nil { return err }\n"+
				"  _ = v\n  return nil\n}\n")
		if d := prv12PreservedUnmarshalErr([]string{q}); len(d) != 0 {
			t.Errorf("guarded unmarshal should be safe, got %+v", d)
		}
	})
}
