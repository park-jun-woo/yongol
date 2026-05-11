//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-09 test — data-component 파일이 존재하지 않는 경우 검증

package stml_openapi

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM09_ComponentNotFound_Positive(t *testing.T) {
	tmpDir := t.TempDir()
	pages := []stml.PageSpec{{
		FileName: "dashboard.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
			Components: []stml.ComponentRef{
				{Name: "NonExistentComponent"},
			},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": getOp("ListUsers", nil, nil),
	})
	fs := makeFS(pages, doc)
	fs.SpecsDir = tmpDir
	diags := Run(fs)
	if !hasDiag(diags, "[TM-09]") {
		t.Errorf("expected TM-09 diagnostic, got %v", diags)
	}
}

func TestTM09_ComponentNotFound_Negative(t *testing.T) {
	tmpDir := t.TempDir()
	compDir := filepath.Join(tmpDir, "frontend", "components")
	if err := os.MkdirAll(compDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(compDir, "DatePicker.tsx"), []byte("export default {}"), 0o644); err != nil {
		t.Fatal(err)
	}

	pages := []stml.PageSpec{{
		FileName: "dashboard.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
			Components: []stml.ComponentRef{
				{Name: "DatePicker"},
			},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": getOp("ListUsers", nil, nil),
	})
	fs := makeFS(pages, doc)
	fs.SpecsDir = tmpDir
	diags := Run(fs)
	if hasDiag(diags, "[TM-09]") {
		t.Errorf("unexpected TM-09 diagnostic, got %v", diags)
	}
}
