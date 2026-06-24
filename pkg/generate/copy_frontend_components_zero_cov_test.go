//ff:func feature=generate type=test control=sequence
//ff:what TestGenerateBatch_ZeroCov — pkg/generate 소형 순수/IO 헬퍼 분기 커버
package generate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFrontendComponents_ZeroCov(t *testing.T) {
	// empty srcFrontendDir → no-op
	if err := copyFrontendComponents("", t.TempDir()); err != nil {
		t.Fatalf("empty srcFrontendDir err: %v", err)
	}
	// missing srcFrontendDir → no-op
	src := filepath.Join(t.TempDir(), "frontend")
	if err := copyFrontendComponents(src, t.TempDir()); err != nil {
		t.Fatalf("missing src err: %v", err)
	}
	// srcFrontendDir is a file not dir → no-op
	fileSrc := filepath.Join(t.TempDir(), "frontend")
	if err := os.WriteFile(fileSrc, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := copyFrontendComponents(fileSrc, t.TempDir()); err != nil {
		t.Fatalf("src-as-file err: %v", err)
	}
	// real srcFrontendDir with a .tsx file → copied into dstSrcDir
	srcRoot := filepath.Join(t.TempDir(), "frontend")
	comp := filepath.Join(srcRoot, "components")
	if err := os.MkdirAll(comp, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(comp, "Foo.tsx"), []byte("export const Foo = 1"), 0o644); err != nil {
		t.Fatal(err)
	}
	dstSrc := filepath.Join(t.TempDir(), "frontend", "src")
	if err := copyFrontendComponents(srcRoot, dstSrc); err != nil {
		t.Fatalf("real copy err: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dstSrc, "components", "Foo.tsx")); err != nil {
		t.Errorf("expected Foo.tsx copied: %v", err)
	}
}
