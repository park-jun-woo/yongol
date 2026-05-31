//ff:func feature=gen-react type=test control=sequence
//ff:what writeIndexHTML index.html 생성 내용·에러경로 검증
package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteIndexHTML(t *testing.T) {
	dir := t.TempDir()
	if err := writeIndexHTML(dir); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "index.html"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "<!DOCTYPE html>")
	assertContains(t, content, `<div id="root"></div>`)
	assertContains(t, content, `<script type="module" src="/src/main.tsx">`)
	assertContains(t, content, "bg-background text-foreground")
}
