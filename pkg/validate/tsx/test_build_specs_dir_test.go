//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=tsx
//ff:what 테스트 헬퍼 — specs/frontend/src 디렉토리 트리와 pages/Home.tsx 스캐폴딩

package tsx

import (
	"os"
	"path/filepath"
	"testing"
)

// buildSpecsDir scaffolds a minimal specs/ tree:
//
//	<tmp>/specs/frontend/src/components/ui/Button.tsx
//	<tmp>/specs/frontend/src/pages/Home.tsx
//
// and returns (specsDir, pageFile).
func buildSpecsDir(t *testing.T, existingComponents []string) (string, string) {
	t.Helper()
	dir := t.TempDir()
	specsDir := filepath.Join(dir, "specs")
	pagesDir := filepath.Join(specsDir, "frontend", "src", "pages")
	if err := os.MkdirAll(pagesDir, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, rel := range existingComponents {
		full := filepath.Join(specsDir, "frontend", "src", rel)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte("export const X = 1;"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	pageFile := filepath.Join(pagesDir, "Home.tsx")
	if err := os.WriteFile(pageFile, []byte("export default function Home() { return null }"), 0o644); err != nil {
		t.Fatal(err)
	}
	return specsDir, pageFile
}
