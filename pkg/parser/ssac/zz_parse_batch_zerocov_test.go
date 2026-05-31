//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestParseBatch_ZeroCov — ssac 파서 헬퍼를 이름으로 직접 호출해 커버 귀속

package ssac

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"
)

const sampleSSaC = `package course

import "context"

type GetCourseRequest struct {
	ID int64
}

// @get Course course = Course.FindByID({ID: request.id})
// @response course
func GetCourse(ctx context.Context, request GetCourseRequest) {}
`

func TestParseSsacStringHelpers_ZeroCov(t *testing.T) {
	// extractInputs / parseInputs
	if _, _, err := extractInputs("{ID: request.id}"); err != nil {
		t.Errorf("extractInputs: %v", err)
	}
	if _, err := parseInputs("{ID: request.id}"); err != nil {
		t.Errorf("parseInputs: %v", err)
	}
	// parseAnnotation dispatch
	if _, err := parseAnnotation("@get Course course = Course.FindByID({ID: request.id})"); err != nil {
		t.Errorf("parseAnnotation get: %v", err)
	}
	// parseCall
	if _, err := parseCall("auth.RefreshRotate({Token: request.token})"); err != nil {
		t.Errorf("parseCall: %v", err)
	}
	// parseEval
	if _, err := parseEval("schedule.IsExpired({At: request.at}) 400 \"past\""); err != nil {
		t.Errorf("parseEval: %v", err)
	}
	// parseCRUDWithResult
	var seq Sequence
	if err := parseCRUDWithResult("Course course = Course.FindByID({ID: request.id})", &seq); err != nil {
		t.Errorf("parseCRUDWithResult: %v", err)
	}
	// parseResponseLine / handleResponseLine
	if _, _, err := parseResponseLine("@response course"); err != nil {
		t.Errorf("parseResponseLine: %v", err)
	}
	handleResponseLine("@response course", nil, false)
	// parseResponseFields
	parseResponseFields([]string{"course: course", "name: instructor.Name"})
}

func TestParseSsacCommentParser_ZeroCov(t *testing.T) {
	cp := &commentParser{}
	if err := cp.processLine("@get Course course = Course.FindByID({ID: request.id})", 1); err != nil {
		t.Errorf("processLine: %v", err)
	}
	cp.inResponse = true
	cp.processResponseBody("course: course")
}

func TestParseSsacASTHelpers_ZeroCov(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "course.ssac", sampleSSaC, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse sample: %v", err)
	}
	imports := collectImports(f)
	_ = imports
	structs := collectStructs(f)
	if len(structs) == 0 {
		t.Error("expected at least one struct")
	}
	for _, decl := range f.Decls {
		collectStructsFromDecl(decl)
	}
	// extractStructInfo via spec
	for _, decl := range f.Decls {
		if gd, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range gd.Specs {
				extractStructInfo(spec)
			}
		}
	}
	// per-func: extractParamInfo, collectFuncComments, parseComments, parseFuncDecl
	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			extractParamInfo(fn)
			collectFuncComments(f, fn.Pos())
			parseFuncDecl(fset, fn, f, "course.ssac", imports, structs)
		}
	}
	// parseComments on the file's comment groups
	for _, cg := range f.Comments {
		parseComments(fset, cg.List)
	}
}

func TestParseDirEntry_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	feature := filepath.Join(dir, "course")
	if err := os.MkdirAll(feature, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(feature, "get_course.ssac")
	if err := os.WriteFile(path, []byte(sampleSSaC), 0o644); err != nil {
		t.Fatal(err)
	}
	sfs, diags := parseDirEntry(dir, path, "get_course.ssac")
	if len(diags) != 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(sfs) == 0 || sfs[0].Feature != "course" {
		t.Fatalf("expected feature course, got %+v", sfs)
	}
	// file directly in dir → error diag
	flat := filepath.Join(dir, "flat.ssac")
	if err := os.WriteFile(flat, []byte(sampleSSaC), 0o644); err != nil {
		t.Fatal(err)
	}
	_, diags2 := parseDirEntry(dir, flat, "flat.ssac")
	if len(diags2) == 0 {
		t.Error("expected error diag for flat file")
	}
}
