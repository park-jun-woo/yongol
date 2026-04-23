//ff:func feature=tsx-parser type=test control=sequence
//ff:what TSX 파서 통합 테스트 — swc @swc/core 가 설치된 환경에서만 실행
package tsx

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// swcAvailable returns true when a @swc/core install is discoverable from
// YONGOL_SWC_PROJECT_DIR or the parent project. Tests skip when absent —
// CI and local dev without Node is expected to skip, not fail, parser tests.
func swcAvailable(t *testing.T) bool {
	t.Helper()
	if _, err := exec.LookPath("node"); err != nil {
		t.Skip("node not available on PATH; skipping tsx parser tests")
		return false
	}
	dir := os.Getenv("YONGOL_SWC_PROJECT_DIR")
	if dir == "" {
		t.Skip("YONGOL_SWC_PROJECT_DIR not set; skipping tsx parser tests (install @swc/core and point YONGOL_SWC_PROJECT_DIR to its parent)")
		return false
	}
	if _, err := os.Stat(filepath.Join(dir, "node_modules", "@swc", "core")); err != nil {
		t.Skipf("@swc/core not installed at %s; skipping", dir)
		return false
	}
	return true
}

func TestParse_SimpleQuery(t *testing.T) {
	if !swcAvailable(t) {
		return
	}
	got, err := Parse("testdata/simple_query.tsx")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(got.Calls) != 1 {
		t.Fatalf("want 1 call, got %d: %+v", len(got.Calls), got.Calls)
	}
	if got.Calls[0].OperationID != "listWorkflows" {
		t.Errorf("operationId: want listWorkflows, got %q", got.Calls[0].OperationID)
	}
	if got.Calls[0].Line == 0 {
		t.Errorf("expected non-zero line")
	}
	if len(got.FormFields) != 0 {
		t.Errorf("want 0 form fields, got %d", len(got.FormFields))
	}
	if len(got.Imports) != 0 {
		t.Errorf("want 0 local imports, got %d: %+v", len(got.Imports), got.Imports)
	}
}

func TestParse_MutationWithForm(t *testing.T) {
	if !swcAvailable(t) {
		return
	}
	got, err := Parse("testdata/mutation_with_form.tsx")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	// apiClient.createWorkflow referenced once (as mutationFn — still a MemberExpression).
	var hasCreate bool
	for _, c := range got.Calls {
		if c.OperationID == "createWorkflow" {
			hasCreate = true
		}
	}
	// Note: `mutationFn: apiClient.createWorkflow` is not a CallExpression;
	// it's a bare reference. Our visitor only picks up invocations.
	// So Calls may be empty — that's correct.
	_ = hasCreate

	wantFields := map[string]bool{"title": true, "trigger_event": true, "description": false}
	gotFields := map[string]bool{}
	for _, f := range got.FormFields {
		gotFields[f.Name] = f.Required
	}
	for name, req := range wantFields {
		r, ok := gotFields[name]
		if !ok {
			t.Errorf("form field %q not found", name)
			continue
		}
		if r != req {
			t.Errorf("form field %q required: want %v, got %v", name, req, r)
		}
	}

	// Component imports: Button, Input from @/components/ui/
	wantImports := map[string]string{"Button": "@/components/ui/Button", "Input": "@/components/ui/Input"}
	for _, imp := range got.Imports {
		if want := wantImports[imp.Name]; want != "" && imp.Path != want {
			t.Errorf("import %s: want path %q, got %q", imp.Name, want, imp.Path)
		}
		delete(wantImports, imp.Name)
	}
	if len(wantImports) != 0 {
		t.Errorf("missing component imports: %v", wantImports)
	}
}

func TestParse_PathParams(t *testing.T) {
	if !swcAvailable(t) {
		return
	}
	got, err := Parse("testdata/path_params.tsx")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(got.Calls) != 1 {
		t.Fatalf("want 1 call, got %d: %+v", len(got.Calls), got.Calls)
	}
	c := got.Calls[0]
	if c.OperationID != "getWorkflow" {
		t.Errorf("operationId: want getWorkflow, got %q", c.OperationID)
	}
	keys := map[string]bool{}
	for _, a := range c.Args {
		keys[a.Key] = true
	}
	if !keys["id"] || !keys["version"] {
		t.Errorf("want args id & version, got %+v", c.Args)
	}
}

func TestParse_NestedCalls(t *testing.T) {
	if !swcAvailable(t) {
		return
	}
	got, err := Parse("testdata/nested_calls.tsx")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	ops := map[string]bool{}
	for _, c := range got.Calls {
		ops[c.OperationID] = true
	}
	if !ops["listWorkflows"] || !ops["listExecutions"] {
		t.Errorf("want listWorkflows + listExecutions, got %v", ops)
	}
}

func TestParse_ComponentImports(t *testing.T) {
	if !swcAvailable(t) {
		return
	}
	got, err := Parse("testdata/component_imports.tsx")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	want := map[string]string{
		"Button":       "@/components/ui/Button",
		"Card":         "@/components/ui/Card",
		"CreditsGauge": "./components/CreditsGauge",
	}
	gotImps := map[string]string{}
	for _, i := range got.Imports {
		gotImps[i.Name] = i.Path
	}
	for name, path := range want {
		if gotImps[name] != path {
			t.Errorf("import %s: want %q, got %q", name, path, gotImps[name])
		}
	}
	// react / clsx must NOT be in Imports (npm packages).
	for _, i := range got.Imports {
		if i.Path == "react" || i.Path == "clsx" {
			t.Errorf("npm package %q should be filtered, got in Imports", i.Path)
		}
	}
}

func TestParse_EmptyPage(t *testing.T) {
	if !swcAvailable(t) {
		return
	}
	got, err := Parse("testdata/empty_page.tsx")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(got.Calls) != 0 || len(got.FormFields) != 0 || len(got.Imports) != 0 {
		t.Errorf("empty page should have nothing, got %+v", got)
	}
}

func TestParseDir(t *testing.T) {
	if !swcAvailable(t) {
		return
	}
	pages, err := ParseDir("testdata")
	if err != nil {
		t.Fatalf("ParseDir: %v", err)
	}
	if len(pages) < 6 {
		t.Errorf("want >=6 pages (fixture count), got %d", len(pages))
	}
}
