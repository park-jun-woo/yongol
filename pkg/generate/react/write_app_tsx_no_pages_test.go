//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX 페이지 없을 때 placeholder 생성 검증

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteAppTSX_NoPages_Placeholder(t *testing.T) {
	dir := t.TempDir()
	if err := writeAppTSX(dir, nil, nil, "", nil, "", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)
	if !strings.Contains(content, "yongol scaffolded frontend") {
		t.Error("expected placeholder content for empty pages")
	}
}
