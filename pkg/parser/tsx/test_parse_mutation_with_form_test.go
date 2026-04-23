//ff:func feature=tsx-parser type=test control=iteration dimension=1
//ff:what Parse — mutation_with_form.tsx 에서 폼 필드 3개 + Button/Input import 추출

package tsx

import "testing"

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
