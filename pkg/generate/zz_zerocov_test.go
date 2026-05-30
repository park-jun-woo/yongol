//ff:func feature=generate type=test control=sequence
//ff:what TestZeroCov — 0% util 함수 (isCopiedExtension / isYongolManaged / mergeFieldlessOps / ResolveBackendType / WithMigration / appendChildNodeFormActions) 회귀

package generate

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestIsCopiedExtension(t *testing.T) {
	for _, p := range []string{"a.tsx", "b.TS", "c.jsx", "d.js", "e.css", "f.svg"} {
		if !isCopiedExtension(p) {
			t.Errorf("isCopiedExtension(%q) = false, want true", p)
		}
	}
	for _, p := range []string{"x.go", "y.json", "z"} {
		if isCopiedExtension(p) {
			t.Errorf("isCopiedExtension(%q) = true, want false", p)
		}
	}
}

func TestIsYongolManaged(t *testing.T) {
	managed := []string{
		"src/api.ts",
		"src/types/Course.ts",
		"src/lib/http.ts",
		"src/components/ui/button.tsx",
	}
	for _, p := range managed {
		if !isYongolManaged(p) {
			t.Errorf("isYongolManaged(%q) = false, want true", p)
		}
	}
	unmanaged := []string{
		"src/pages/Home.tsx",
		"src/components/MyCard.tsx",
		"README.md",
	}
	for _, p := range unmanaged {
		if isYongolManaged(p) {
			t.Errorf("isYongolManaged(%q) = true, want false", p)
		}
	}
}

func TestMergeFieldlessOps(t *testing.T) {
	dst := map[string]bool{"A": true}
	src := map[string]bool{"B": true, "C": true}
	mergeFieldlessOps(dst, src)
	for _, k := range []string{"A", "B", "C"} {
		if !dst[k] {
			t.Errorf("dst missing %q after merge", k)
		}
	}
}

func TestResolveBackendType(t *testing.T) {
	cases := []struct {
		lang, fw string
		want     BackendType
		err      bool
	}{
		{"go", "gin", GoGin, false},
		{"typescript", "nestjs", NestJS, false},
		{"python", "fastapi", FastAPI, false},
		{"ruby", "rails", "", true},
	}
	for _, c := range cases {
		got, err := ResolveBackendType(c.lang, c.fw)
		if c.err {
			if err == nil {
				t.Errorf("ResolveBackendType(%q,%q) expected error", c.lang, c.fw)
			}
			continue
		}
		if err != nil || got != c.want {
			t.Errorf("ResolveBackendType(%q,%q) = (%q,%v), want %q", c.lang, c.fw, got, err, c.want)
		}
	}
}

func TestWithMigration(t *testing.T) {
	var buf bytes.Buffer
	opt := WithMigration(MigrationHook{Version: "v1", Logger: &buf})
	var c generateConfig
	opt(&c)
	if c.migration.Version != "v1" {
		t.Errorf("WithMigration did not set version, got %q", c.migration.Version)
	}
}

func TestAppendChildNodeFormActions(t *testing.T) {
	seen := map[string]bool{}
	// action branch with fields
	res := appendChildNodeFormActions(nil, stmlparser.ChildNode{
		Kind:   "action",
		Action: &stmlparser.ActionBlock{OperationID: "CreateX", Fields: []stmlparser.FieldBind{{Name: "n"}}},
	}, seen)
	if len(res) != 1 || res[0].opID != "CreateX" {
		t.Fatalf("action branch = %+v", res)
	}
	// duplicate operationId is skipped
	res = appendChildNodeFormActions(res, stmlparser.ChildNode{
		Kind:   "action",
		Action: &stmlparser.ActionBlock{OperationID: "CreateX", Fields: []stmlparser.FieldBind{{Name: "n"}}},
	}, seen)
	if len(res) != 1 {
		t.Errorf("duplicate not skipped: %+v", res)
	}
	// recurse branches: fetch / state / static / each with nested action
	nested := stmlparser.ChildNode{
		Kind:   "action",
		Action: &stmlparser.ActionBlock{OperationID: "Nested", Fields: []stmlparser.FieldBind{{Name: "m"}}},
	}
	branches := []stmlparser.ChildNode{
		{Kind: "fetch", Fetch: &stmlparser.FetchBlock{Children: []stmlparser.ChildNode{nested}}},
		{Kind: "state", State: &stmlparser.StateBind{Children: []stmlparser.ChildNode{nested}}},
		{Kind: "static", Static: &stmlparser.StaticElement{Children: []stmlparser.ChildNode{nested}}},
		{Kind: "each", Each: &stmlparser.EachBlock{Children: []stmlparser.ChildNode{nested}}},
	}
	for _, br := range branches {
		out := appendChildNodeFormActions(nil, br, map[string]bool{})
		if len(out) != 1 || out[0].opID != "Nested" {
			t.Errorf("recurse branch %q = %+v", br.Kind, out)
		}
	}
	// unknown kind → no change
	out := appendChildNodeFormActions(nil, stmlparser.ChildNode{Kind: "bind"}, map[string]bool{})
	if len(out) != 0 {
		t.Errorf("bind kind should add nothing: %+v", out)
	}
}

func TestCopyUserComponentFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "src.txt")
	dst := filepath.Join(dir, "out", "dst.txt")
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(src, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := copyUserComponentFile(src, dst); err != nil {
		t.Fatalf("copy: %v", err)
	}
	got, err := os.ReadFile(dst)
	if err != nil || string(got) != "hello" {
		t.Errorf("dst content = %q err=%v", got, err)
	}

	// open error: missing src.
	if err := copyUserComponentFile(filepath.Join(dir, "nope"), dst); err == nil {
		t.Error("expected open error")
	}
	// create error: dst dir does not exist.
	if err := copyUserComponentFile(src, filepath.Join(dir, "missing-dir", "x.txt")); err == nil {
		t.Error("expected create error")
	}
}

func TestLogIncrementalMigration_ZeroCov(t *testing.T) {
	var buf bytes.Buffer
	res := &migration.Result{
		MigrationFile: "0002_x.up.sql",
		OpsCount:      1,
		Operations:    []migration.Operation{migration.DropCheck{Table: "t", Name: "c"}},
	}
	logIncrementalMigration(MigrationHook{Logger: &buf}, res)
	s := buf.String()
	if !strings.Contains(s, "incremental") || !strings.Contains(s, "0002_x.up.sql") || !strings.Contains(s, "drop check c") {
		t.Errorf("log output = %q", s)
	}
}

func TestLogMigrationWarnings_ZeroCov(t *testing.T) {
	// nil logger → no panic, early return.
	logMigrationWarnings(MigrationHook{Logger: nil}, []diagnostic.Diagnostic{{Level: diagnostic.LevelWarning, Message: "w"}})

	var buf bytes.Buffer
	diags := []diagnostic.Diagnostic{
		{Level: diagnostic.LevelWarning, Message: "warn1"},
		{Level: diagnostic.LevelError, Message: "err1"},
	}
	logMigrationWarnings(MigrationHook{Logger: &buf}, diags)
	s := buf.String()
	if !strings.Contains(s, "warn1") {
		t.Errorf("missing warning: %q", s)
	}
	if strings.Contains(s, "err1") {
		t.Errorf("error should not be logged: %q", s)
	}
}

func TestMakeFrontendCopyWalker_ZeroCov(t *testing.T) {
	srcRoot := t.TempDir()
	dstRoot := t.TempDir()
	// create: a copied file, a managed file (skipped), a non-copied ext (skipped),
	// node_modules dir (skipped).
	mustWrite := func(rel, content string) string {
		p := filepath.Join(srcRoot, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}
	plain := mustWrite("src/pages/Home.tsx", "x")
	managed := mustWrite("src/api.ts", "y")
	other := mustWrite("src/readme.md", "z")
	nm := mustWrite("node_modules/pkg/index.js", "n")

	walker := makeFrontendCopyWalker(srcRoot, dstRoot)

	// walkErr propagation.
	if err := walker("p", nil, errTestWalk); err != errTestWalk {
		t.Errorf("walkErr not propagated: %v", err)
	}

	// node_modules dir → SkipDir.
	nmDir := filepath.Join(srcRoot, "node_modules")
	fi, _ := os.Stat(nmDir)
	if err := walker(nmDir, fi, nil); err != filepath.SkipDir {
		t.Errorf("node_modules should SkipDir, got %v", err)
	}
	// non-node_modules dir → nil.
	pagesFi, _ := os.Stat(filepath.Join(srcRoot, "src", "pages"))
	if err := walker(filepath.Join(srcRoot, "src", "pages"), pagesFi, nil); err != nil {
		t.Errorf("dir walk = %v", err)
	}

	walkFile := func(p string) error {
		fi, err := os.Stat(p)
		if err != nil {
			t.Fatal(err)
		}
		return walker(p, fi, nil)
	}
	if err := walkFile(plain); err != nil {
		t.Errorf("plain copy err: %v", err)
	}
	if err := walkFile(managed); err != nil {
		t.Errorf("managed err: %v", err)
	}
	if err := walkFile(other); err != nil {
		t.Errorf("other err: %v", err)
	}
	_ = nm

	// plain should be copied; managed and other should not.
	if _, err := os.Stat(filepath.Join(dstRoot, "src/pages/Home.tsx")); err != nil {
		t.Errorf("plain not copied: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dstRoot, "src/api.ts")); !os.IsNotExist(err) {
		t.Error("managed file should not be copied")
	}
	if _, err := os.Stat(filepath.Join(dstRoot, "src/readme.md")); !os.IsNotExist(err) {
		t.Error("non-copied ext should not be copied")
	}
}

var errTestWalk = errors.New("walk failure")
