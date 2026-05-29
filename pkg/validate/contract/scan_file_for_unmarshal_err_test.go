//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanFileForUnmarshalErr — Unmarshal 호출 에러 미처리 수집 검증

package contract

import (
	"path/filepath"
	"testing"
)

func TestScanFileForUnmarshalErr(t *testing.T) {
	dir := t.TempDir()

	t.Run("ignored unmarshal error flagged", func(t *testing.T) {
		p := filepath.Join(dir, "u.go")
		writePreserved(t, p,
			"package service\nfunc F(b []byte) {\n  var v T\n  json.Unmarshal(b, &v)\n  _ = v\n}\n")
		if d := scanFileForUnmarshalErr(p); len(d) != 1 {
			t.Fatalf("expected 1 diag, got %+v", d)
		}
	})

	t.Run("explicit discard → safe", func(t *testing.T) {
		p := filepath.Join(dir, "ok.go")
		writePreserved(t, p,
			"package service\nfunc G(b []byte) {\n  var v T\n  _ = json.Unmarshal(b, &v)\n  _ = v\n}\n")
		if d := scanFileForUnmarshalErr(p); len(d) != 0 {
			t.Errorf("expected no diag, got %+v", d)
		}
	})
}
