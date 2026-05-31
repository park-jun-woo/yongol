//ff:func feature=gen-react type=test control=sequence
//ff:what writeViteConfig vite.config.ts 생성 내용·에러경로 검증
package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteViteConfig(t *testing.T) {
	dir := t.TempDir()
	if err := writeViteConfig(dir); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "vite.config.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import react from '@vitejs/plugin-react'")
	assertContains(t, content, "plugins: [react()]")
	assertContains(t, content, "'@': path.resolve(__dirname, 'src')")
	assertContains(t, content, "'/api': 'http://localhost:8080'")
}
