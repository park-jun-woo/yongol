//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWriteMainGo — cmd/main.go 조립 + 빈 라인/import 필터/error 경로 검증

package boot

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteMainGo(t *testing.T) {
	imports := []string{`"os"`, `"time"`}
	body := []string{
		`_ = os.Getenv("X")`, // references os -> kept
		"",                   // blank line -> emitted as bare newline
		`r := gin.Default()`,
	}

	t.Run("WritesMainGo", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeMainGo(dir, imports, body); err != nil {
			t.Fatalf("writeMainGo error: %v", err)
		}
		raw, err := os.ReadFile(filepath.Join(dir, "backend", "cmd", "main.go"))
		if err != nil {
			t.Fatalf("read main.go: %v", err)
		}
		got := string(raw)
		if !strings.Contains(got, "func main() {") {
			t.Errorf("missing main func:\n%s", got)
		}
		if !strings.Contains(got, `"os"`) {
			t.Errorf("expected used import 'os':\n%s", got)
		}
		if strings.Contains(got, `"time"`) {
			t.Errorf("unused import 'time' should be filtered:\n%s", got)
		}
		if !strings.Contains(got, "\t_ = os.Getenv") {
			t.Errorf("expected indented body line:\n%s", got)
		}
	})

	t.Run("MkdirError", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "backend"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeMainGo(dir, imports, body); err == nil {
			t.Errorf("expected MkdirAll error, got nil")
		}
	})

	t.Run("WriteFileError", func(t *testing.T) {
		dir := t.TempDir()
		cmdDir := filepath.Join(dir, "backend", "cmd")
		if err := os.MkdirAll(filepath.Join(cmdDir, "main.go"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeMainGo(dir, imports, body); err == nil {
			t.Errorf("expected WriteFile error, got nil")
		}
	})
}
