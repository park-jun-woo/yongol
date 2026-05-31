//ff:func feature=gen-react type=test control=sequence
//ff:what resolveOpenapiTsBinary env→PATH→npx 해결 순서·부재 에러 검증
package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveOpenapiTsBinaryFromPATH(t *testing.T) {
	isolatePATH(t)

	// Place an executable named openapi-typescript on PATH.
	pathDir := t.TempDir()
	binPath := filepath.Join(pathDir, "openapi-typescript")
	if err := os.WriteFile(binPath, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", pathDir)

	argv, _, err := resolveOpenapiTsBinary()
	if err != nil {
		t.Fatal(err)
	}
	if len(argv) != 1 || argv[0] != binPath {
		t.Errorf("expected resolved PATH binary [%s], got %v", binPath, argv)
	}
}
