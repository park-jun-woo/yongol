//ff:func feature=gen-react type=test control=sequence
//ff:what writeLibArtifacts — lib 디렉토리 생성 실패 시 writeLibUtils 에러 전파 검증
package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteLibArtifacts_MkdirError(t *testing.T) {
	// srcDir is a regular file, so writeLibUtils' MkdirAll(srcDir/lib)
	// must fail (a path component is not a directory) and the error must
	// propagate out of writeLibArtifacts before breadcrumbs are emitted.
	file := filepath.Join(t.TempDir(), "notadir")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := writeLibArtifacts(file, nil, nil); err == nil {
		t.Fatal("expected error creating lib dir under a file, got nil")
	}
}
