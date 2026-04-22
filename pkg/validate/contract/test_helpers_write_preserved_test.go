//ff:func feature=validate-contract type=test-helper control=sequence
//ff:what writePreserved — 테스트용 preserved 파일 작성 헬퍼 (stale //ff:checked hash 삽입)

package contract

import (
	"os"
	"testing"
)

// writePreserved writes a file with a stale //ff:checked hash so
// DetectPreserved classifies it as StatePreserved during tests.
func writePreserved(t *testing.T, path, body string) {
	t.Helper()
	src := "//ff:func feature=service type=handler control=sequence\n" +
		"//ff:what stub\n" +
		"//ff:checked llm=yongol-gen hash=deadbeef\n" +
		body
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
}
