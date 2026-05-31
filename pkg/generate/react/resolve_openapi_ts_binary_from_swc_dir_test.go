//ff:func feature=gen-react type=test control=sequence
//ff:what resolveOpenapiTsBinary env→PATH→npx 해결 순서·부재 에러 검증
package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveOpenapiTsBinaryFromSwcDir(t *testing.T) {
	isolatePATH(t)

	swcDir := t.TempDir()
	binDir := filepath.Join(swcDir, "node_modules", ".bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatal(err)
	}
	binPath := filepath.Join(binDir, "openapi-typescript")
	if err := os.WriteFile(binPath, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("YONGOL_SWC_PROJECT_DIR", swcDir)

	argv, _, err := resolveOpenapiTsBinary()
	if err != nil {
		t.Fatal(err)
	}
	if len(argv) != 1 || argv[0] != binPath {
		t.Errorf("expected argv [%s], got %v", binPath, argv)
	}
}
