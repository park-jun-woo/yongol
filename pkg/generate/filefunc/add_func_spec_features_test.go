//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestAddFuncSpecFeatures — Func 스펙 package→description 맵 추가/보존 검증

package filefunc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestAddFuncSpecFeatures(t *testing.T) {
	fs := &yongol.Fullstack{
		ProjectFuncSpecs: []funcspec.FuncSpec{
			{Package: "auth", Description: "authentication"},
			{Package: "  billing  ", Description: "  payments  "},
			{Package: "auth", Description: "second-auth"}, // existing non-empty desc preserved
			{Package: "", Description: "ignored"},         // empty package skipped
		},
	}
	dst := map[string]string{}
	addFuncSpecFeatures(dst, fs)

	if got := dst["auth"]; got != "authentication" {
		t.Errorf("auth: expected existing desc preserved, got %q", got)
	}
	if got := dst["billing"]; got != "payments" {
		t.Errorf("billing: expected trimmed desc, got %q", got)
	}
	if _, ok := dst[""]; ok {
		t.Errorf("empty package should not be inserted")
	}
	if len(dst) != 2 {
		t.Errorf("expected 2 entries, got %d: %v", len(dst), dst)
	}
}
