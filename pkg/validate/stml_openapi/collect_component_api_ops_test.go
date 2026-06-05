//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestCollectComponentApiOpsIntersection — .tsx api.X에서 실재 op만 교집합 채택

package stml_openapi

import "testing"

func TestCollectComponentApiOpsIntersection(t *testing.T) {
	specsDir := t.TempDir()
	// DatePicker calls a real op (api.LoadDates) and a non-existent one (api.Bogus).
	writeComponent(t, specsDir, "DatePicker", `export function DatePicker() {
	api.LoadDates();
	api.Bogus();
	return null;
}`)

	names := map[string]struct{}{"DatePicker": {}}
	ops := map[string]struct{}{"LoadDates": {}, "Other": {}}
	out := make(map[string]struct{})

	collectComponentApiOps(names, specsDir, ops, out)

	if _, ok := out["LoadDates"]; !ok {
		t.Error("expected real op LoadDates to be consumed")
	}
	if _, ok := out["Bogus"]; ok {
		t.Error("non-existent op Bogus must not be consumed (intersection only)")
	}
	if len(out) != 1 {
		t.Errorf("expected exactly 1 consumed op, got %+v", out)
	}
}
