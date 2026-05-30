//ff:func feature=gen-splitter type=test
//ff:what zz_zerocov_test — splitter 패키지의 0% 커버리지 함수(cleanOriginal/preserveComments/isPreservedFile/SplitDirectory) 단위 테스트
package splitter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"
)

func writeFileZeroCov(t *testing.T, dir, name, content string) string {
	t.Helper()
	p := filepath.Join(dir, name)
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", name, err)
	}
	return p
}

func TestIsPreservedFile_ZeroCov(t *testing.T) {
	cases := []struct {
		name string
		tool Tool
		want bool
	}{
		{"querier.go", ToolSQLC, true},
		{"db.go", ToolSQLC, true},
		{"models.go", ToolSQLC, false},
		{"foo.sql.go", ToolSQLC, false},
		{"querier.go", ToolOAPICodegen, false},
		{"anything.go", Tool("unknown"), false},
		// basename is taken even with a path prefix.
		{filepath.Join("sub", "querier.go"), ToolSQLC, true},
	}
	for _, c := range cases {
		if got := isPreservedFile(c.name, c.tool); got != c.want {
			t.Errorf("isPreservedFile(%q,%q)=%v want %v", c.name, c.tool, got, c.want)
		}
	}
}

func TestPreserveComments_ZeroCov(t *testing.T) {
	src := `package p

// doc comment
//go:generate something
func Foo() {}

func Bar() {}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "x.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	cmap := ast.NewCommentMap(fset, file, file.Comments)

	var fooDecl, barDecl ast.Decl
	for _, d := range file.Decls {
		if fd, ok := d.(*ast.FuncDecl); ok {
			switch fd.Name.Name {
			case "Foo":
				fooDecl = d
			case "Bar":
				barDecl = d
			}
		}
	}
	if fooDecl == nil || barDecl == nil {
		t.Fatal("decls not found")
	}

	groups := preserveComments(cmap, fooDecl)
	if len(groups) == 0 {
		t.Fatal("expected comment groups for Foo")
	}
	for _, g := range groups {
		if g == nil {
			t.Fatal("nil group leaked into output")
		}
	}

	// Bar has no associated comments → empty slice.
	if got := preserveComments(cmap, barDecl); len(got) != 0 {
		t.Errorf("expected no comments for Bar, got %d", len(got))
	}
}

func TestCleanOriginal_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	// sqlc tool: models.go and foo.sql.go are originals; querier.go preserved;
	// split.go in keep; nested dir is skipped.
	writeFileZeroCov(t, dir, "models.go", "package p\n")
	writeFileZeroCov(t, dir, "foo.sql.go", "package p\n")
	writeFileZeroCov(t, dir, "querier.go", "package p\n")
	writeFileZeroCov(t, dir, "user.model.go", "package p\n") // a split result, in keep
	writeFileZeroCov(t, dir, "unrelated.txt", "x\n")
	if err := os.Mkdir(filepath.Join(dir, "nested"), 0o755); err != nil {
		t.Fatal(err)
	}

	keep := map[string]bool{"user.model.go": true}
	if err := cleanOriginal(dir, ToolSQLC, keep); err != nil {
		t.Fatalf("cleanOriginal: %v", err)
	}

	mustGone := []string{"models.go", "foo.sql.go"}
	for _, n := range mustGone {
		if _, err := os.Stat(filepath.Join(dir, n)); !os.IsNotExist(err) {
			t.Errorf("%s should have been removed", n)
		}
	}
	mustStay := []string{"querier.go", "user.model.go", "unrelated.txt"}
	for _, n := range mustStay {
		if _, err := os.Stat(filepath.Join(dir, n)); err != nil {
			t.Errorf("%s should have survived: %v", n, err)
		}
	}
}

func TestCleanOriginal_ReadDirError_ZeroCov(t *testing.T) {
	if err := cleanOriginal(filepath.Join(t.TempDir(), "does-not-exist"), ToolSQLC, nil); err == nil {
		t.Fatal("expected error for missing dir")
	}
}

func TestSplitDirectory_NoDir_ZeroCov(t *testing.T) {
	// Non-existent directory → no-op, nil error.
	if err := SplitDirectory(filepath.Join(t.TempDir(), "absent"), ToolSQLC); err != nil {
		t.Fatalf("expected nil for absent dir, got %v", err)
	}
}

func TestSplitDirectory_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	// A real sqlc-style source file to split, plus a preserved querier.go,
	// plus a nested dir to exercise the IsDir skip.
	writeFileZeroCov(t, dir, "models.go", "package db\n\ntype User struct {\n\tID int64\n}\n")
	writeFileZeroCov(t, dir, "querier.go", "package db\n\ntype Querier interface{}\n")
	if err := os.Mkdir(filepath.Join(dir, "sub"), 0o755); err != nil {
		t.Fatal(err)
	}

	if err := SplitDirectory(dir, ToolSQLC); err != nil {
		t.Fatalf("SplitDirectory: %v", err)
	}

	// models.go should be split away (removed) and querier.go preserved.
	if _, err := os.Stat(filepath.Join(dir, "models.go")); !os.IsNotExist(err) {
		t.Error("models.go should have been split and removed")
	}
	if _, err := os.Stat(filepath.Join(dir, "querier.go")); err != nil {
		t.Errorf("querier.go should be preserved: %v", err)
	}
}

func TestSplitDirectory_SplitFileError_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	// A models.go that fails to parse → SplitFile returns an error which
	// SplitDirectory wraps and returns.
	writeFileZeroCov(t, dir, "models.go", "package db\nthis is not valid go\n")
	if err := SplitDirectory(dir, ToolSQLC); err == nil {
		t.Fatal("expected split error to propagate")
	}
}
