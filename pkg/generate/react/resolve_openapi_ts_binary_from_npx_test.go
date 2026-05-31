//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what resolveOpenapiTsBinary env→PATH→npx 해결 순서·부재 에러 검증
package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveOpenapiTsBinaryFromNpx(t *testing.T) {
	isolatePATH(t)

	// Only npx is available on PATH.
	pathDir := t.TempDir()
	npxPath := filepath.Join(pathDir, "npx")
	if err := os.WriteFile(npxPath, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", pathDir)

	argv, _, err := resolveOpenapiTsBinary()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"npx", "--yes", "openapi-typescript"}
	if len(argv) != len(want) {
		t.Fatalf("expected %v, got %v", want, argv)
	}
	for i := range want {
		if argv[i] != want[i] {
			t.Errorf("argv[%d] = %q, want %q", i, argv[i], want[i])
		}
	}
}
