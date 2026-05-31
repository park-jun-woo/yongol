//ff:func feature=generate type=test control=sequence
//ff:what TestGenerateBatch_ZeroCov — pkg/generate 소형 순수/IO 헬퍼 분기 커버
package generate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFrontendComponents_ZeroCov(t *testing.T) {
	// empty specsDir → no-op
	if err := copyFrontendComponents("", t.TempDir()); err != nil {
		t.Fatalf("empty specsDir err: %v", err)
	}
	// missing frontend dir → no-op
	specs := t.TempDir()
	if err := copyFrontendComponents(specs, t.TempDir()); err != nil {
		t.Fatalf("missing frontend err: %v", err)
	}
	// frontend is a file not dir → no-op
	if err := os.WriteFile(filepath.Join(specs, "frontend"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := copyFrontendComponents(specs, t.TempDir()); err != nil {
		t.Fatalf("frontend-as-file err: %v", err)
	}
	// real frontend dir with a .tsx file → copied
	specs2 := t.TempDir()
	fe := filepath.Join(specs2, "frontend", "components")
	if err := os.MkdirAll(fe, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(fe, "Foo.tsx"), []byte("export const Foo = 1"), 0o644); err != nil {
		t.Fatal(err)
	}
	arts := t.TempDir()
	if err := copyFrontendComponents(specs2, arts); err != nil {
		t.Fatalf("real copy err: %v", err)
	}
	if _, err := os.Stat(filepath.Join(arts, "frontend", "src", "components", "Foo.tsx")); err != nil {
		t.Errorf("expected Foo.tsx copied: %v", err)
	}
}
