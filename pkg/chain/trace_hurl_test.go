//ff:func feature=chain type=test control=sequence
//ff:what traceHurl 가 endpoint 참조 .hurl 파일을 찾고 nil 조건(doc nil/미매칭)을 처리하는지 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestTraceHurl(t *testing.T) {
	specsDir := t.TempDir()
	testsDir := filepath.Join(specsDir, "tests")
	if err := os.MkdirAll(testsDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	hurl := "GET {{base}}/courses/42\nHTTP 200\n"
	if err := os.WriteFile(filepath.Join(testsDir, "get_course.hurl"), []byte(hurl), 0o644); err != nil {
		t.Fatalf("write hurl: %v", err)
	}

	paths := openapi3.NewPaths()
	paths.Set("/courses/42", &openapi3.PathItem{
		Get: &openapi3.Operation{OperationID: "GetCourse"},
	})
	doc := &openapi3.T{Paths: paths}

	links := traceHurl("GetCourse", doc, testsDir, specsDir)
	if len(links) != 1 {
		t.Fatalf("expected 1 hurl link, got %d", len(links))
	}
	if links[0].Kind != "Hurl" || links[0].File != "tests/get_course.hurl" {
		t.Errorf("link fields: %+v", links[0])
	}
	if links[0].Line != 1 {
		t.Errorf("line: got %d, want 1", links[0].Line)
	}
	if links[0].Summary != "scenario: get_course.hurl" {
		t.Errorf("summary: %q", links[0].Summary)
	}

	// doc nil → nil.
	if traceHurl("GetCourse", nil, testsDir, specsDir) != nil {
		t.Error("expected nil for nil doc")
	}
	// operationId not in paths → nil.
	if traceHurl("Unknown", doc, testsDir, specsDir) != nil {
		t.Error("expected nil for unknown operationId")
	}
}
