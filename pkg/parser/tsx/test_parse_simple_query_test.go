//ff:func feature=tsx-parser type=test control=sequence
//ff:what Parse — simple_query.tsx 에서 apiClient.listWorkflows 1개 호출 추출

package tsx

import "testing"

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
