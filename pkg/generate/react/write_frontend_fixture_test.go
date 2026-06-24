//ff:func feature=gen-react type=test control=sequence
//ff:what writeFrontendFixture — tsc 게이트 테스트 픽스처(frontend/tsconfig+main.ts, deps 시 node_modules 마커) 생성 헬퍼 (BUG-137 Phase041)

package react

import (
	"os"
	"path/filepath"
	"testing"
)

// writeFrontendFixture lays out <dir>/frontend with a tsconfig and a single
// source file, returning the frontend dir itself (the arg RunTscCheck now
// takes directly). When deps is true an empty node_modules marker is created so
// RunTscCheck attempts the gate.
func writeFrontendFixture(t *testing.T, src string, deps bool) string {
	t.Helper()
	artifacts := t.TempDir()
	fe := filepath.Join(artifacts, "frontend")
	if err := os.MkdirAll(fe, 0o755); err != nil {
		t.Fatal(err)
	}
	tsconfig := `{ "compilerOptions": { "strict": true, "noEmit": true, "skipLibCheck": true } }`
	if err := os.WriteFile(filepath.Join(fe, "tsconfig.json"), []byte(tsconfig), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(fe, "main.ts"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	if deps {
		if err := os.MkdirAll(filepath.Join(fe, "node_modules"), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	return fe
}
