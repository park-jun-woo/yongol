//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestAddServiceFuncFeatures — SSaC ServiceFunc.Feature 맵 추가 검증

package filefunc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestAddServiceFuncFeatures(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{Feature: "auth"},
			{Feature: "  workflow  "},
			{Feature: "auth"}, // duplicate, ignored
			{Feature: ""},     // empty, skipped
		},
	}
	dst := map[string]string{"auth": "existing"} // existing key not overwritten
	addServiceFuncFeatures(dst, fs)

	if got := dst["auth"]; got != "existing" {
		t.Errorf("auth: expected existing value preserved, got %q", got)
	}
	if _, ok := dst["workflow"]; !ok {
		t.Errorf("expected trimmed workflow key inserted")
	}
	if dst["workflow"] != "" {
		t.Errorf("workflow desc should be empty, got %q", dst["workflow"])
	}
	if _, ok := dst[""]; ok {
		t.Errorf("empty feature should not be inserted")
	}
	if len(dst) != 2 {
		t.Errorf("expected 2 entries, got %d: %v", len(dst), dst)
	}
}
