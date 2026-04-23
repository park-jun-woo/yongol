//ff:func feature=tsx-parser type=test control=iteration dimension=1
//ff:what Parse — nested_calls.tsx 에서 listWorkflows + listExecutions 동시 추출

package tsx

import "testing"

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
