//ff:func feature=gen-react type=test control=sequence
//ff:what writeTailwindConfig index.css 생성 검증

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteTailwindConfig_IndexCSSEmitted(t *testing.T) {
	dir := t.TempDir()
	srcDir := filepath.Join(dir, "src")
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := writeTailwindConfig(dir, nil, nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(srcDir, "index.css"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)
	if !strings.Contains(content, "@tailwind base") {
		t.Error("expected @tailwind base in index.css")
	}
}
