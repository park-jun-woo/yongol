//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestWriteMethodFile_ZeroCov — method Go 파일 기록 + import 정렬/dedup
package ssac

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteMethodFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	err := writeMethodFile(
		dir,
		"GetWidget",
		"example.com/app",
		[]string{`"context"`, `"context"`, ``, `"log/slog"`},
		"func (server *Server) GetWidget(ctx context.Context) error",
		[]string{"slog.Info(\"hi\")", "", "return nil"},
		"위젯 조회",
	)
	if err != nil {
		t.Fatalf("writeMethodFile: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "get_widget.go"))
	if err != nil {
		t.Fatalf("read emitted file: %v", err)
	}
	src := string(data)
	for _, want := range []string{
		"package service",
		`"context"`,
		`"log/slog"`,
		"func (server *Server) GetWidget(ctx context.Context) error {",
		"return nil",
	} {
		if !strings.Contains(src, want) {
			t.Errorf("missing %q in:\n%s", want, src)
		}
	}
	// dedup: "context" must appear exactly once in the import lines.
	if strings.Count(src, "\t\"context\"\n") != 1 {
		t.Errorf("expected single context import, src:\n%s", src)
	}
}
