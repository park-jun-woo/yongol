//ff:func feature=tsx-parser type=test control=iteration dimension=1
//ff:what Parse — path_params.tsx 에서 getWorkflow + id/version 인자 키 추출

package tsx

import "testing"

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
