//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what zz_zerocov_writers — 0% 생성/기록 헬퍼(generateServerHelpers/writeMethodFile/writeConvertFunc/writeConvertListFunc/emitConvert*File) 검증

package ssac

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

//ff:what TestGenerateServerGo_ZeroCov — internal/service/server.go (Server struct) 생성
func TestGenerateServerGo_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := generateServerGo(&yongol.Fullstack{}, dir, "example.com/app"); err != nil {
		t.Fatalf("generateServerGo: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "backend", "internal", "service", "server.go"))
	if err != nil {
		t.Fatalf("read server.go: %v", err)
	}
	src := string(data)
	for _, want := range []string{
		"package service",
		"type Server struct {",
		"DB      *pgxpool.Pool",
		"Queries *db.Queries",
		`"example.com/app/internal/db"`,
	} {
		if !strings.Contains(src, want) {
			t.Errorf("missing %q in:\n%s", want, src)
		}
	}
}

// convertSchemaZeroCov returns a small api schema with one required scalar
// and one optional scalar property for the convert* renderers.
func convertSchemaZeroCov() *openapi3.Schema {
	s := openapi3.NewSchema()
	s.Type = &openapi3.Types{"object"}
	s.Required = []string{"id"}
	s.Properties = openapi3.Schemas{
		"id":   &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		"name": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
	}
	return s
}

//ff:what TestWriteConvertFunc_ZeroCov — convert<Name> 본문 (required/optional 분기)
func TestWriteConvertFunc_ZeroCov(t *testing.T) {
	var sb strings.Builder
	writeConvertFunc(&sb, "Widget", convertSchemaZeroCov(), nil)
	out := sb.String()
	for _, want := range []string{
		"func convertWidget(row db.Widget) (*api.Widget, error) {",
		"return &api.Widget{",
		"Id:",
		"Name:",
		"}, nil",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}
}

//ff:what TestWriteConvertListFunc_ZeroCov — convert<Name>List 본문
func TestWriteConvertListFunc_ZeroCov(t *testing.T) {
	var sb strings.Builder
	writeConvertListFunc(&sb, "Widget")
	out := sb.String()
	for _, want := range []string{
		"func convertWidgetList(rows []db.Widget) ([]api.Widget, error) {",
		"result := make([]api.Widget, len(rows))",
		"item, err := convertWidget(row)",
		"return result, nil",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}
}

//ff:what TestGenerateServerHelpers_ZeroCov — ptr_of/deref_* 헬퍼 파일 emit
func TestGenerateServerHelpers_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := generateServerHelpers(dir); err != nil {
		t.Fatalf("generateServerHelpers: %v", err)
	}
	serviceDir := filepath.Join(dir, "backend", "internal", "service")
	entries, err := os.ReadDir(serviceDir)
	if err != nil {
		t.Fatalf("read service dir: %v", err)
	}
	if len(entries) == 0 {
		t.Fatalf("expected helper files emitted")
	}
	foundPtrOf := false
	for _, e := range entries {
		if e.Name() == "ptr_of.go" {
			foundPtrOf = true
		}
	}
	if !foundPtrOf {
		t.Errorf("expected ptr_of.go, got %v", entries)
	}
}

//ff:what TestWriteMethodFile_ZeroCov — method Go 파일 기록 + import 정렬/dedup
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

//ff:what TestEmitConvertFiles_ZeroCov — emitConvertFuncFile / emitConvertListFile 파일 기록
func TestEmitConvertFiles_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	used := map[string]bool{}

	if err := emitConvertFuncFile(dir, "example.com/app", "Widget", convertSchemaZeroCov(), nil, used); err != nil {
		t.Fatalf("emitConvertFuncFile: %v", err)
	}
	if err := emitConvertListFile(dir, "example.com/app", "Widget", used); err != nil {
		t.Fatalf("emitConvertListFile: %v", err)
	}

	got, _ := os.ReadDir(dir)
	if len(got) == 0 {
		t.Fatalf("expected converter files emitted, got none")
	}
	// at least one file references convertWidget
	found := false
	for _, e := range got {
		b, _ := os.ReadFile(filepath.Join(dir, e.Name()))
		if strings.Contains(string(b), "convertWidget") {
			found = true
		}
	}
	if !found {
		t.Errorf("expected a file containing convertWidget, files=%v", got)
	}
}
