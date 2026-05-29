//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestCollectFileImports — 파일 import 목록에서 패키지 이름 집합 추출 검증

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCollectFileImports(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "imports.go")
	src := "package p\n\n" +
		"import (\n" +
		"\t\"database/sql\"\n" +
		"\t\"fmt\"\n" +
		"\tsqlx \"github.com/jmoiron/sqlx\"\n" +
		")\n"
	if err := os.WriteFile(p, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}

	pkgs := collectFileImports(p)
	for _, want := range []string{"sql", "fmt", "sqlx"} {
		if !pkgs[want] {
			t.Fatalf("expected import %q in %v", want, pkgs)
		}
	}

	t.Run("unparseable file yields empty", func(t *testing.T) {
		bad := filepath.Join(dir, "bad.go")
		if err := os.WriteFile(bad, []byte("not go at all {{{"), 0o644); err != nil {
			t.Fatal(err)
		}
		if got := collectFileImports(bad); len(got) != 0 {
			t.Fatalf("expected empty set, got %v", got)
		}
	})
}
